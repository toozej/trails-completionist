package generator

import "embed"

//go:embed *.html.tmpl
var Templates embed.FS
