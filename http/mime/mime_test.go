package mime

import (
	"testing"
)

func TestType_String(t *testing.T) {
	cases := map[string]struct {
		mime Type
		want string
	}{
		"application_octet_stream": {mime: ApplicationOctetStream, want: "application/octet-stream"},
		"application_json":         {mime: ApplicationJSON, want: "application/json"},
		"application_xml":          {mime: ApplicationXML, want: "application/xml"},
		"text_plain":               {mime: TextPlain, want: "text/plain"},
		"text_html":                {mime: TextHTML, want: "text/html"},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := tc.mime.String()
			if g, w := got, tc.want; g != w {
				t.Errorf("string mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
