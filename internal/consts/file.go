package consts

import "golang.org/x/tools/imports"

const (
	PARAM_NAME = "BuilderGen"

	DEFAULT_BUILDER_SUFFIX = "_builder.go"

	BUILD_HEADER  = "// Code generated by " + PARAM_NAME + " "
	BUILD_PACKAGE = "package"
)

var ImportOptions = &imports.Options{
	FormatOnly: false,
	TabIndent:  true,
	Comments:   true,
}
