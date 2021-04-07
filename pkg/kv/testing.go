package kv

import (
	"bytes"
	"errors"
	"io/ioutil"
	"strings"
	"testing"
)

type NewStoreFunc func() (Store, error)

func TestStorePushPullSequence(t *testing.T, s NewStoreFunc) {
	cases := map[string]struct {
		key  string
		data string
	}{
		"default": {
			key:  "abc",
			data: "xyz",
		},
		"empty_data": {
			key:  "abc",
			data: "", // empty
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup store
			store, err := s()
			if err != nil {
				t.Fatalf("failed to create store: %v", err)
			}

			// Assert data not exist
			if err := store.Pull(tc.key, ioutil.Discard); !errors.Is(err, ErrNotFound) {
				t.Fatalf("data exists: %v", err)
			}

			// Push data
			if err := store.Push(tc.key, strings.NewReader(tc.data)); err != nil {
				t.Fatalf("failed to push data: %v", err)
			}

			// Pull data
			var b bytes.Buffer
			if err := store.Pull(tc.key, &b); err != nil {
				t.Fatalf("failed to pull data: %v", err)
			}

			// Assert data match
			if g, w := b.String(), tc.data; g != w {
				t.Errorf("data mismatch:\ngot:\t%#v\nwant:\t%#v", g, w)
			}
		})
	}
}

func TestStorePushError(t *testing.T, s NewStoreFunc) {
	cases := map[string]struct {
		key  string
		data string
		err  error
	}{
		"error_empty_key": {
			key:  "", // empty
			data: "xyz",
			err:  ErrEmptyKey,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup store
			store, err := s()
			if err != nil {
				t.Fatalf("failed to create store: %v", err)
			}

			// Call method
			if err := store.Push(tc.key, strings.NewReader(tc.data)); !errors.Is(err, tc.err) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", errors.Unwrap(err), tc.err)
			}
		})
	}
}

func TestStorePullError(t *testing.T, s NewStoreFunc) {
	cases := map[string]struct {
		key  string
		data string
		err  error
	}{
		"error_empty_key": {
			key:  "", // empty
			data: "xyz",
			err:  ErrEmptyKey,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup store
			store, err := s()
			if err != nil {
				t.Fatalf("failed to create store: %v", err)
			}

			// Call method
			if err := store.Pull(tc.key, ioutil.Discard); !errors.Is(err, tc.err) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", errors.Unwrap(err), tc.err)
			}
		})
	}
}

func TestStoreClear(t *testing.T, s NewStoreFunc) {
	cases := map[string]struct {
		key  string
		data string
	}{
		"default": {
			key:  "abc",
			data: "xyz",
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup store
			store, err := s()
			if err != nil {
				t.Fatalf("failed to create store: %v", err)
			}

			// Assert data not exist
			if err := store.Pull(tc.key, ioutil.Discard); !errors.Is(err, ErrNotFound) {
				t.Fatalf("data exists: %v", err)
			}

			// Push data
			if err := store.Push(tc.key, strings.NewReader(tc.data)); err != nil {
				t.Fatalf("failed to push data: %v", err)
			}

			// Assert data exists
			if err := store.Pull(tc.key, ioutil.Discard); errors.Is(err, ErrNotFound) {
				t.Fatalf("data not exists: %v", err)
			}

			// Clear data
			if err := store.Clear(tc.key); err != nil {
				t.Fatalf("failed to clear data: %v", err)
			}

			// Assert data not exist
			if err := store.Pull(tc.key, ioutil.Discard); !errors.Is(err, ErrNotFound) {
				t.Fatalf("data exists: %v", err)
			}
		})
	}
}

func TestStoreClearError(t *testing.T, s NewStoreFunc) {
	cases := map[string]struct {
		key  string
		data string
		err  error
	}{
		"error_empty_key": {
			key:  "", // empty
			data: "xyz",
			err:  ErrEmptyKey,
		},
	}
	for name, tc := range cases {
		tc := tc // capture range variable
		t.Run(name, func(t *testing.T) {
			t.Parallel()

			// Setup store
			store, err := s()
			if err != nil {
				t.Fatalf("failed to create store: %v", err)
			}

			// Call method
			if err := store.Clear(tc.key); !errors.Is(err, tc.err) {
				t.Fatalf("error value mismatch:\ngot:\t%#v\nwant:\t%#v", errors.Unwrap(err), tc.err)
			}
		})
	}
}
