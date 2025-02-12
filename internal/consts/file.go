package consts

const (
	PARAM_NAME = "BuilderGen"
	VERSION    = "v0.3.0"

	DEFAULT_BUILDER_SUFFIX = "_builder.go"

	BUILD_HEADER  = "// Code generated by " + PARAM_NAME + " " + VERSION
	BUILD_PACKAGE = "package"

	DEFAULT_TRIM = "\n\r\t "

	DEFAULT_TEMP_DIR = "."
)
