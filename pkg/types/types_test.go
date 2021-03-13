package types

import (
	"bytes"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"testing"
)

func TestNullUUID_Value(t *testing.T) {
	cases := map[string]struct {
		obj  driver.Valuer
		want []byte
		err  error
	}{
		"string_default": {
			obj: NullUUID{
				String: "abc",
				Valid:  true,
			},
			want: []byte("abc"),
		},
		"string_empty": {
			obj: NullUUID{
				String: "",
				Valid:  true,
			},
			want: []byte(""),
		},
		"string_null": {
			obj: NullUUID{
				String: "",
				Valid:  false,
			},
			want: []byte(nil), // should be nil
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := tc.obj.Value()
			if g, w := err, tc.err; !errors.Is(g, w) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			switch v := got.(type) {
			case nil:
				if g, w := v, tc.want; g != nil {
					t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
				}
			case []byte:
				if g, w := v, tc.want; !bytes.Equal(g, w) {
					t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
				}
			default:
				t.Fatalf("unhandled type:\ngot:\t%#v", v)
			}
		})
	}
}

func TestNullUUID_Scan(t *testing.T) {
	cases := map[string]struct {
		input interface{}
		want  NullUUID
		err   error
	}{
		"string_default": {
			input: "abc",
			want: NullUUID{
				String: "abc",
				Valid:  true,
			},
		},
		"string_empty": {
			input: "",
			want: NullUUID{
				String: "",
				Valid:  true,
			},
		},
		"string_null": {
			input: nil,
			want: NullUUID{
				String: "",
				Valid:  false, // should be nil
			},
		},
		"bytestring_default": {
			input: []byte("abc"),
			want: NullUUID{
				String: "abc",
				Valid:  true,
			},
		},
		"bytestring_empty": {
			input: []byte(""),
			want: NullUUID{
				String: "",
				Valid:  false,
			},
		},
		"bytestring_null": {
			input: []byte(nil),
			want: NullUUID{
				String: "",
				Valid:  false, // should be nil
			},
		},
		"error_unhandled_type_int": {
			input: int(0),
			want: NullUUID{
				String: "",
				Valid:  false, // should be nil
			},
			err: ErrUnhandledType,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var obj NullUUID
			err := obj.Scan(tc.input)
			if g, w := err, tc.err; !errors.Is(g, w) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := obj.String, tc.want.String; g != w {
				t.Errorf("field String mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := obj.Valid, tc.want.Valid; g != w {
				t.Errorf("field Valid mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestNullUUID_MarshalJSON(t *testing.T) {
	cases := map[string]struct {
		obj  json.Marshaler
		want []byte
		err  error
	}{
		"string_default": {
			obj: NullUUID{
				String: "abc",
				Valid:  true,
			},
			want: []byte("abc"),
		},
		"string_empty": {
			obj: NullUUID{
				String: "",
				Valid:  true,
			},
			want: []byte(""),
		},
		"string_null": {
			obj: NullUUID{
				String: "",
				Valid:  false,
			},
			want: []byte(nil), // should be nil
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got, err := tc.obj.MarshalJSON()
			if g, w := err, tc.err; !errors.Is(g, w) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := got, tc.want; !bytes.Equal(g, w) {
				t.Errorf("value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestNullUUID_UnmarshalJSON(t *testing.T) {
	cases := map[string]struct {
		input []byte
		want  NullUUID
		err   error
	}{
		"string_default": {
			input: []byte("abc"),
			want: NullUUID{
				String: "abc",
				Valid:  true,
			},
		},
		"string_empty": {
			input: []byte(""),
			want: NullUUID{
				String: "",
				Valid:  true,
			},
		},
		"string_null": {
			input: []byte(nil),
			want: NullUUID{
				String: "",
				Valid:  false, // should be nil
			},
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			var obj NullUUID
			err := obj.UnmarshalJSON(tc.input)
			if g, w := err, tc.err; !errors.Is(g, w) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := obj.String, tc.want.String; g != w {
				t.Errorf("field String mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := obj.Valid, tc.want.Valid; g != w {
				t.Errorf("field Valid mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
