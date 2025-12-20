package consts

import (
	"errors"

	"golang.org/x/tools/imports"
)

const (
	ErrMsgSrcNotfound          = "source file is required"
	ErrMsgNoStructsFound       = "source file has no structs"
	ErrMsgNameNotFound         = "name is required"
	ErrMsgTypeNotFound         = "type not found for field"
	ErrMsgNotGoFile            = "source is not a valid go file"
	ErrMsgInvalidStruct        = "invalid struct type"
	ErrMsgInvalidConfigFile    = "invalid config file"
	ErrMsgSyntax               = "syntax error"
	ErrMsgPkgNotFound          = "package not found"
	ErrMsgTargetStructNotFound = "target struct is not found"
	ErrMsgDone                 = "done"
	ErrMsgInvalidParserMode    = "invalid parser mode"
)

var (
	ErrSrcNotFound       = errors.New(ErrMsgSrcNotfound)
	ErrNoStructsFound    = errors.New(ErrMsgNoStructsFound)
	ErrNameNotFound      = errors.New(ErrMsgNameNotFound)
	ErrTypeNotfound      = errors.New(ErrMsgTypeNotFound)
	ErrNotGoFile         = errors.New(ErrMsgNotGoFile)
	ErrInvalidStructType = errors.New(ErrMsgInvalidStruct)
	ErrSyntaxErr         = errors.New(ErrMsgSyntax)
	ErrPackageNotFound   = errors.New(ErrMsgPkgNotFound)
	ErrNotFound          = errors.New(ErrMsgTargetStructNotFound)
	ErrDone              = errors.New(ErrMsgDone)
	ErrInvalidConfigFile = errors.New(ErrMsgInvalidConfigFile)
	ErrInvalidParserMode = errors.New(ErrMsgInvalidParserMode)

	ImportOptions = &imports.Options{
		FormatOnly: false,
		TabIndent:  true,
		Comments:   true,
	}
)
