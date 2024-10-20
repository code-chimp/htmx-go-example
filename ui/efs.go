package ui

import "embed"

// Files is an embedded filesystem that includes the "html" templates.
//
//go:embed "html"
var Files embed.FS
