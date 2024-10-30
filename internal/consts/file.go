package consts

const (
	PARAM_NAME = "BuilderGen"
	VERSION    = "v0.2.1"

	DEFAULT_BUILDER_SUFFIX = "_builder.go"

	BUILD_HEADER  = "// Code generated by " + PARAM_NAME + " " + VERSION
	BUILD_PACKAGE = "package"

	DEFAULT_TRIM = "\n\r\t "
)
