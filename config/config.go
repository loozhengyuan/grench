// Package config provides the tools for managing application configuration.
package config

import (
	"fmt"
	"os"
	"path/filepath"
)

// Node is the interface that all configuration nodes must implement.
type Node interface{}

// ResolvePath returns the canonical path to the configuration
// file.
//
// ResolvePath tries to first return the absolute path of the
// configuration file relative to the user's configuration
// directory using os.UserConfigDir. Else, it tries to return
// the absolute path of the configuration file relative to the
// current working directory using os.Getwd. Else, it returns
// the relative path to the configuration file.
func ResolvePath(elem ...string) string {
	p := filepath.Join(elem...)
	if d, err := os.UserConfigDir(); err == nil { // if no err
		return filepath.Join(d, p)
	}
	if d, err := os.Getwd(); err == nil { // if no err
		return filepath.Join(d, p)
	}
	return p
}

// TouchFile creates a configuration file with mode 0600 if it
// does not already exist.
//
// If the base directory does not exist, it will be created
// with mode 0700.
func TouchFile(name string) error {
	if err := os.MkdirAll(filepath.Dir(name), 0700); err != nil {
		return fmt.Errorf("create app dir: %w", err)
	}
	f, err := os.OpenFile(name, os.O_RDONLY|os.O_CREATE, 0600)
	if err != nil {
		return fmt.Errorf("create config file: %w", err)
	}
	defer f.Close()
	if err := f.Close(); err != nil {
		return fmt.Errorf("close file: %w", err)
	}
	return nil
}
