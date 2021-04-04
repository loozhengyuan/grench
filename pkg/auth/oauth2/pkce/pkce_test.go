package pkce

import (
	"testing"
)

func TestCodeVerifier(t *testing.T) {
	cases := map[string]struct {
		src                []byte
		codeChallengePlain string
		codeChallengeS256  string
		err                error
	}{
		"rfc7636_reference_implementation": {
			src: []byte{
				116, 24, 223, 180, 151, 153, 224, 37, 79, 250, 96, 125, 216, 173,
				187, 186, 22, 212, 37, 77, 105, 214, 191, 240, 91, 88, 5, 88, 83,
				132, 141, 121,
			},
			codeChallengePlain: "dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
			codeChallengeS256:  "E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM",
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Create new code verifier
			c := NewCodeVerifier(tc.src)

			// Assert plain code_challenge (i.e. code_verifier value)
			if g, w := c.Value(), tc.codeChallengePlain; g != w {
				t.Errorf("code challenge plain mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}

			// Assert sha256 code_challenge (i.e. code_verifier value)
			if g, w := c.CodeChallengeSHA256(), tc.codeChallengeS256; g != w {
				t.Errorf("code challenge sha256 mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}
