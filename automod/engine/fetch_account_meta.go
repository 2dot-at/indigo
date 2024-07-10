package engine

import (
	"context"
	"encoding/json"
	"fmt"

	comatproto "github.com/bluesky-social/indigo/api/atproto"
	appbsky "github.com/bluesky-social/indigo/api/bsky"
	toolsozone "github.com/bluesky-social/indigo/api/ozone"
	"github.com/bluesky-social/indigo/atproto/identity"
	"github.com/bluesky-social/indigo/atproto/syntax"
)

// Helper to hydrate metadata about an account from several sources: PDS (if access), mod service (if access), public identity resolution
func (e *Engine) GetAccountMeta(ctx context.Context, ident *identity.Identity) (*AccountMeta, error) {

	logger := e.Logger.With("did", ident.DID.String())

	// fallback in case client wasn't configured (eg, testing)
	if e.BskyClient == nil {
		logger.Warn("skipping account meta hydration")
		am := AccountMeta{
			Identity: ident,
			Profile:  ProfileSummary{},
		}
		return &am, nil
	}

	existing, err := e.Cache.Get(ctx, "acct", ident.DID.String())
	if err != nil {
		return nil, fmt.Errorf("failed checking account meta cache: %w", err)
	}
	if existing != "" {
		var am AccountMeta
		err := json.Unmarshal([]byte(existing), &am)
		if err != nil {
			return nil, fmt.Errorf("parsing AccountMeta from cache: %v", err)
		}
		am.Identity = ident
		return &am, nil
	}

	// doing a "full" fetch from here on
	accountMetaFetches.Inc()
	am := AccountMeta{
		Identity: ident,
	}

	// automod-internal "flags"
	flags, err := e.Flags.Get(ctx, ident.DID.String())
	if err != nil {
		return nil, fmt.Errorf("failed checking account flag cache: %w", err)
	}
	am.AccountFlags = flags

	// fetch account metadata from AppView
	pv, err := appbsky.ActorGetProfile(ctx, e.BskyClient, ident.DID.String())
	if err != nil {
		logger.Warn("account profile lookup failed", "err", err)
		return &am, nil
	}

	am.Profile = ProfileSummary{
		HasAvatar:   pv.Avatar != nil,
		Description: pv.Description,
		DisplayName: pv.DisplayName,
	}
	if pv.PostsCount != nil {
		am.PostsCount = *pv.PostsCount
	}
	if pv.FollowersCount != nil {
		am.FollowersCount = *pv.FollowersCount
	}
	if pv.FollowsCount != nil {
		am.FollowsCount = *pv.FollowsCount
	}

	var labels []string
	var negLabels []string
	for _, lbl := range pv.Labels {
		if lbl.Neg != nil && *lbl.Neg == true {
			negLabels = append(negLabels, lbl.Val)
		} else {
			labels = append(labels, lbl.Val)
		}
	}
	am.AccountLabels = dedupeStrings(labels)
	am.AccountNegatedLabels = dedupeStrings(negLabels)

	if pv.CreatedAt != nil {
		ts, err := syntax.ParseDatetimeTime(*pv.CreatedAt)
		if err != nil {
			logger.Warn("invalid profile createdAt", "err", err, "createdAt", pv.CreatedAt)
		} else {
			am.CreatedAt = &ts
		}
	}

	// first attempt to fetch private account metadata from Ozone
	if e.OzoneClient != nil {
		rd, err := toolsozone.ModerationGetRepo(ctx, e.OzoneClient, ident.DID.String())
		if err != nil {
			logger.Warn("failed to fetch private account metadata from Ozone", "err", err)
		} else {
			ap := AccountPrivate{}
			if rd.Email != nil && *rd.Email != "" {
				ap.Email = *rd.Email
			}
			if rd.EmailConfirmedAt != nil && *rd.EmailConfirmedAt != "" {
				ap.EmailConfirmed = true
			}
			ts, err := syntax.ParseDatetimeTime(rd.IndexedAt)
			if err != nil {
				return nil, fmt.Errorf("bad account IndexedAt: %w", err)
			}
			ap.IndexedAt = ts
			if rd.DeactivatedAt != nil {
				am.Deactivated = true
			}
			if rd.Moderation != nil && rd.Moderation.SubjectStatus != nil {
				if rd.Moderation.SubjectStatus.Takendown != nil && *rd.Moderation.SubjectStatus.Takendown == true {
					am.Takendown = true
				}
				ap.AccountTags = dedupeStrings(rd.Moderation.SubjectStatus.Tags)
			}
			am.Private = &ap
		}
	} else if e.AdminClient != nil {
		// fall back to PDS/entryway fetching; less metadata available
		pv, err := comatproto.AdminGetAccountInfo(ctx, e.AdminClient, ident.DID.String())
		if err != nil {
			logger.Warn("failed to fetch private account metadata from PDS/entryway", "err", err)
		} else {
			ap := AccountPrivate{}
			if pv.Email != nil && *pv.Email != "" {
				ap.Email = *pv.Email
			}
			if pv.EmailConfirmedAt != nil && *pv.EmailConfirmedAt != "" {
				ap.EmailConfirmed = true
			}
			ts, err := syntax.ParseDatetimeTime(pv.IndexedAt)
			if err != nil {
				return nil, fmt.Errorf("bad account IndexedAt: %w", err)
			}
			ap.IndexedAt = ts
			am.Private = &ap
		}
	}

	val, err := json.Marshal(&am)
	if err != nil {
		return nil, err
	}

	if err := e.Cache.Set(ctx, "acct", ident.DID.String(), string(val)); err != nil {
		logger.Error("writing to account meta cache failed", "err", err)
	}
	return &am, nil
}
