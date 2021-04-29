package config

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var (
	testUserConfigDir  string
	testUserWorkingDir string
)

func run() error {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return fmt.Errorf("get user config dir: %w", err)
	}
	testUserConfigDir = configDir
	workingDir, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("get user working dir: %w", err)
	}
	testUserWorkingDir = workingDir
	return nil
}

func TestMain(m *testing.M) {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestResolvePath(t *testing.T) {
	cases := map[string]struct {
		input []string
		want  string
	}{
		"user_config_dir": {
			input: []string{"path", "to", "file.txt"},
			want:  filepath.Join(testUserConfigDir, "path", "to", "file.txt"),
		},
		// TODO: working dir
		// TODO: relative path
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			if got := ResolvePath(tc.input...); got != tc.want {
				t.Errorf("filepath mismatch:\ngot:\t%#v\nwant:\t%#v", got, tc.want)
			}
		})
	}
}

func TestTouchFile(t *testing.T) {
	cases := map[string]struct {
		path string
	}{
		"file_exists": {
			path: filepath.Join("path", "to", "file.txt"),
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup temp dir
			tmp, err := ioutil.TempDir("", "")
			if err != nil {
				t.Fatalf("failed to create temp dir: %v", err)
			}
			defer os.RemoveAll(tmp)

			// Join file path
			p := filepath.Join(tmp, tc.path)

			// Assert file not exists
			if _, err := os.Stat(p); !errors.Is(err, os.ErrNotExist) {
				t.Fatalf("error mismatch:\ngot:\t%#v\nwant:\t%#v", err, os.ErrNotExist)
			}

			// Execute function
			if err := TouchFile(p); err != nil {
				t.Fatalf("failed to execute function: %v", err)
			}

			// Stat file
			fi, err := os.Stat(p)
			if err != nil {
				t.Fatalf("failed to stat file: %v", err)
			}

			// Assert file attributes
			if g, w := fi.IsDir(), false; g != w {
				t.Errorf("file type mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
			if g, w := fi.Mode().Perm(), os.FileMode(0600); g != w {
				t.Errorf("file mode mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}

			// Write to file
			data := []byte("something")
			if err := ioutil.WriteFile(p, data, 0600); err != nil {
				t.Fatalf("failed to write to file: %v", err)
			}

			// Execute function again (assert idempotence)
			if err := TouchFile(p); err != nil {
				t.Fatalf("failed to execute function: %v", err)
			}

			// Assert file content unmodified
			f, err := os.Open(p)
			if err != nil {
				t.Fatalf("failed to open file: %v", err)
			}
			b, err := ioutil.ReadAll(f)
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			if g, w := b, data; !bytes.Equal(g, w) {
				t.Errorf("file contents mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}

			// TODO: Assert directory permissions
		})
	}
}
