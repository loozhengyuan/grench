// Package mime provides the constants for MIME types.
package mime

// Type is the type definition of a MIME type.
type Type string

// String returns the string representation of the type.
func (t Type) String() string {
	return string(t)
}

// Common MIME types as defined by Mozilla.
// https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types
const (
	ApplicationOctetStream Type = "application/octet-stream"
	ApplicationJSON        Type = "application/json"
	ApplicationXML         Type = "application/xml"

	TextPlain Type = "text/plain"
	TextHTML  Type = "text/html"
)
