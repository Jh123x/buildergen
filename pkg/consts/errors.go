package consts

import (
	"golang.org/x/tools/imports"
)

var (
	ImportOptions = &imports.Options{
		FormatOnly: false,
		TabIndent:  true,
		Comments:   true,
	}
)
