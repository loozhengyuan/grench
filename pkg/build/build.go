// Package build provides the build information of an application.
package build

import (
	"encoding/json"
	"fmt"
	"io"
)

// Info represents the build information of the application.
type Info struct {
	App       string `json:"app"`
	System    string `json:"system"`
	Arch      string `json:"architecture"`
	Version   string `json:"version"`
	Commit    string `json:"commit"`
	Timestamp string `json:"timestamp"`
}

// OutputText outputs the build info to io.Writer as an unstructured,
// human-readable text format.
func (i Info) OutputText(w io.Writer) error {
	if _, err := fmt.Fprintf(w, "%s %s %s/%s %s %s\n", i.App, i.Version, i.System, i.Arch, i.Timestamp, i.Commit); err != nil {
		return err
	}
	return nil
}

// OutputJSON outputs the build info to io.Writer as a structured JSON
// format.
func (i Info) OutputJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(i)
}
