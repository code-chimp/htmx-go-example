package ui

import "embed"

// Files is an embedded filesystem that includes the "html" and "static" directories.
//
//go:embed "html" "static"
var Files embed.FS
