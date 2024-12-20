// Code generated by github.com/whyrusleeping/cbor-gen. DO NOT EDIT.

package labels

import (
	"fmt"
	"io"
	"math"
	"sort"

	cid "github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	xerrors "golang.org/x/xerrors"
)

var _ = xerrors.Errorf
var _ = cid.Undef
var _ = math.E
var _ = sort.Sort

func (t *UnsignedLabel) MarshalCBOR(w io.Writer) error {
	if t == nil {
		_, err := w.Write(cbg.CborNull)
		return err
	}

	cw := cbg.NewCborWriter(w)
	fieldCount := 8

	if t.Cid == nil {
		fieldCount--
	}

	if t.Exp == nil {
		fieldCount--
	}

	if t.Neg == nil {
		fieldCount--
	}

	if t.Ver == nil {
		fieldCount--
	}

	if _, err := cw.Write(cbg.CborEncodeMajorType(cbg.MajMap, uint64(fieldCount))); err != nil {
		return err
	}

	// t.Cid (string) (string)
	if t.Cid != nil {

		if len("cid") > 1000000 {
			return xerrors.Errorf("Value in field \"cid\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("cid"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("cid")); err != nil {
			return err
		}

		if t.Cid == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if len(*t.Cid) > 1000000 {
				return xerrors.Errorf("Value in field t.Cid was too long")
			}

			if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.Cid))); err != nil {
				return err
			}
			if _, err := cw.WriteString(string(*t.Cid)); err != nil {
				return err
			}
		}
	}

	// t.Cts (string) (string)
	if len("cts") > 1000000 {
		return xerrors.Errorf("Value in field \"cts\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("cts"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("cts")); err != nil {
		return err
	}

	if len(t.Cts) > 1000000 {
		return xerrors.Errorf("Value in field t.Cts was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Cts))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Cts)); err != nil {
		return err
	}

	// t.Exp (string) (string)
	if t.Exp != nil {

		if len("exp") > 1000000 {
			return xerrors.Errorf("Value in field \"exp\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("exp"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("exp")); err != nil {
			return err
		}

		if t.Exp == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if len(*t.Exp) > 1000000 {
				return xerrors.Errorf("Value in field t.Exp was too long")
			}

			if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(*t.Exp))); err != nil {
				return err
			}
			if _, err := cw.WriteString(string(*t.Exp)); err != nil {
				return err
			}
		}
	}

	// t.Neg (bool) (bool)
	if t.Neg != nil {

		if len("neg") > 1000000 {
			return xerrors.Errorf("Value in field \"neg\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("neg"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("neg")); err != nil {
			return err
		}

		if t.Neg == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if err := cbg.WriteBool(w, *t.Neg); err != nil {
				return err
			}
		}
	}

	// t.Src (string) (string)
	if len("src") > 1000000 {
		return xerrors.Errorf("Value in field \"src\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("src"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("src")); err != nil {
		return err
	}

	if len(t.Src) > 1000000 {
		return xerrors.Errorf("Value in field t.Src was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Src))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Src)); err != nil {
		return err
	}

	// t.Uri (string) (string)
	if len("uri") > 1000000 {
		return xerrors.Errorf("Value in field \"uri\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("uri"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("uri")); err != nil {
		return err
	}

	if len(t.Uri) > 1000000 {
		return xerrors.Errorf("Value in field t.Uri was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Uri))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Uri)); err != nil {
		return err
	}

	// t.Val (string) (string)
	if len("val") > 1000000 {
		return xerrors.Errorf("Value in field \"val\" was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("val"))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string("val")); err != nil {
		return err
	}

	if len(t.Val) > 1000000 {
		return xerrors.Errorf("Value in field t.Val was too long")
	}

	if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len(t.Val))); err != nil {
		return err
	}
	if _, err := cw.WriteString(string(t.Val)); err != nil {
		return err
	}

	// t.Ver (int64) (int64)
	if t.Ver != nil {

		if len("ver") > 1000000 {
			return xerrors.Errorf("Value in field \"ver\" was too long")
		}

		if err := cw.WriteMajorTypeHeader(cbg.MajTextString, uint64(len("ver"))); err != nil {
			return err
		}
		if _, err := cw.WriteString(string("ver")); err != nil {
			return err
		}

		if t.Ver == nil {
			if _, err := cw.Write(cbg.CborNull); err != nil {
				return err
			}
		} else {
			if *t.Ver >= 0 {
				if err := cw.WriteMajorTypeHeader(cbg.MajUnsignedInt, uint64(*t.Ver)); err != nil {
					return err
				}
			} else {
				if err := cw.WriteMajorTypeHeader(cbg.MajNegativeInt, uint64(-*t.Ver-1)); err != nil {
					return err
				}
			}
		}

	}
	return nil
}

func (t *UnsignedLabel) UnmarshalCBOR(r io.Reader) (err error) {
	*t = UnsignedLabel{}

	cr := cbg.NewCborReader(r)

	maj, extra, err := cr.ReadHeader()
	if err != nil {
		return err
	}
	defer func() {
		if err == io.EOF {
			err = io.ErrUnexpectedEOF
		}
	}()

	if maj != cbg.MajMap {
		return fmt.Errorf("cbor input should be of type map")
	}

	if extra > cbg.MaxLength {
		return fmt.Errorf("UnsignedLabel: map struct too large (%d)", extra)
	}

	n := extra

	nameBuf := make([]byte, 3)
	for i := uint64(0); i < n; i++ {
		nameLen, ok, err := cbg.ReadFullStringIntoBuf(cr, nameBuf, 1000000)
		if err != nil {
			return err
		}

		if !ok {
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(cr, func(cid.Cid) {}); err != nil {
				return err
			}
			continue
		}

		switch string(nameBuf[:nameLen]) {
		// t.Cid (string) (string)
		case "cid":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadStringWithMax(cr, 1000000)
					if err != nil {
						return err
					}

					t.Cid = (*string)(&sval)
				}
			}
			// t.Cts (string) (string)
		case "cts":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Cts = string(sval)
			}
			// t.Exp (string) (string)
		case "exp":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					sval, err := cbg.ReadStringWithMax(cr, 1000000)
					if err != nil {
						return err
					}

					t.Exp = (*string)(&sval)
				}
			}
			// t.Neg (bool) (bool)
		case "neg":

			{
				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}

					maj, extra, err = cr.ReadHeader()
					if err != nil {
						return err
					}
					if maj != cbg.MajOther {
						return fmt.Errorf("booleans must be major type 7")
					}

					var val bool
					switch extra {
					case 20:
						val = false
					case 21:
						val = true
					default:
						return fmt.Errorf("booleans are either major type 7, value 20 or 21 (got %d)", extra)
					}
					t.Neg = &val
				}
			}
			// t.Src (string) (string)
		case "src":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Src = string(sval)
			}
			// t.Uri (string) (string)
		case "uri":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Uri = string(sval)
			}
			// t.Val (string) (string)
		case "val":

			{
				sval, err := cbg.ReadStringWithMax(cr, 1000000)
				if err != nil {
					return err
				}

				t.Val = string(sval)
			}
			// t.Ver (int64) (int64)
		case "ver":
			{

				b, err := cr.ReadByte()
				if err != nil {
					return err
				}
				if b != cbg.CborNull[0] {
					if err := cr.UnreadByte(); err != nil {
						return err
					}
					maj, extra, err := cr.ReadHeader()
					if err != nil {
						return err
					}
					var extraI int64
					switch maj {
					case cbg.MajUnsignedInt:
						extraI = int64(extra)
						if extraI < 0 {
							return fmt.Errorf("int64 positive overflow")
						}
					case cbg.MajNegativeInt:
						extraI = int64(extra)
						if extraI < 0 {
							return fmt.Errorf("int64 negative overflow")
						}
						extraI = -1 - extraI
					default:
						return fmt.Errorf("wrong type for int64 field: %d", maj)
					}

					t.Ver = (*int64)(&extraI)
				}
			}

		default:
			// Field doesn't exist on this type, so ignore it
			if err := cbg.ScanForLinks(r, func(cid.Cid) {}); err != nil {
				return err
			}
		}
	}

	return nil
}
