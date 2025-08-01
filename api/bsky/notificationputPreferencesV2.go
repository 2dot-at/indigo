// Code generated by cmd/lexgen (see Makefile's lexgen); DO NOT EDIT.

package bsky

// schema: app.bsky.notification.putPreferencesV2

import (
	"context"

	"github.com/bluesky-social/indigo/lex/util"
)

// NotificationPutPreferencesV2_Input is the input argument to a app.bsky.notification.putPreferencesV2 call.
type NotificationPutPreferencesV2_Input struct {
	Chat              *NotificationDefs_ChatPreference       `json:"chat,omitempty" cborgen:"chat,omitempty"`
	Follow            *NotificationDefs_FilterablePreference `json:"follow,omitempty" cborgen:"follow,omitempty"`
	Like              *NotificationDefs_FilterablePreference `json:"like,omitempty" cborgen:"like,omitempty"`
	LikeViaRepost     *NotificationDefs_FilterablePreference `json:"likeViaRepost,omitempty" cborgen:"likeViaRepost,omitempty"`
	Mention           *NotificationDefs_FilterablePreference `json:"mention,omitempty" cborgen:"mention,omitempty"`
	Quote             *NotificationDefs_FilterablePreference `json:"quote,omitempty" cborgen:"quote,omitempty"`
	Reply             *NotificationDefs_FilterablePreference `json:"reply,omitempty" cborgen:"reply,omitempty"`
	Repost            *NotificationDefs_FilterablePreference `json:"repost,omitempty" cborgen:"repost,omitempty"`
	RepostViaRepost   *NotificationDefs_FilterablePreference `json:"repostViaRepost,omitempty" cborgen:"repostViaRepost,omitempty"`
	StarterpackJoined *NotificationDefs_Preference           `json:"starterpackJoined,omitempty" cborgen:"starterpackJoined,omitempty"`
	SubscribedPost    *NotificationDefs_Preference           `json:"subscribedPost,omitempty" cborgen:"subscribedPost,omitempty"`
	Unverified        *NotificationDefs_Preference           `json:"unverified,omitempty" cborgen:"unverified,omitempty"`
	Verified          *NotificationDefs_Preference           `json:"verified,omitempty" cborgen:"verified,omitempty"`
}

// NotificationPutPreferencesV2_Output is the output of a app.bsky.notification.putPreferencesV2 call.
type NotificationPutPreferencesV2_Output struct {
	Preferences *NotificationDefs_Preferences `json:"preferences" cborgen:"preferences"`
}

// NotificationPutPreferencesV2 calls the XRPC method "app.bsky.notification.putPreferencesV2".
func NotificationPutPreferencesV2(ctx context.Context, c util.LexClient, input *NotificationPutPreferencesV2_Input) (*NotificationPutPreferencesV2_Output, error) {
	var out NotificationPutPreferencesV2_Output
	if err := c.LexDo(ctx, util.Procedure, "application/json", "app.bsky.notification.putPreferencesV2", nil, input, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
