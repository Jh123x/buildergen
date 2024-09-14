package consts

import (
	"errors"

	"golang.org/x/tools/imports"
)

const (
	ErrMsgSrcNotfound       = "source file is required"
	ErrMsgNoStructsFound    = "source file has no structs"
	ErrMsgNameNotFound      = "name is required"
	ErrMsgTypeNotFound      = "type not found for field"
	ErrMsgNotGoFile         = "source is not a valid go file"
	ErrMsgInvalidStruct     = "invalid struct type"
	ErrMsgInvalidConfigFile = "invalid config file"
)

var (
	ErrSrcNotFound       = errors.New(ErrMsgSrcNotfound)
	ErrNoStructsFound    = errors.New(ErrMsgNoStructsFound)
	ErrNameNotFound      = errors.New(ErrMsgNameNotFound)
	ErrTypeNotfound      = errors.New(ErrMsgTypeNotFound)
	ErrNotGoFile         = errors.New(ErrMsgNotGoFile)
	ErrInvalidStructType = errors.New(ErrMsgInvalidStruct)
	ErrSyntaxErr         = errors.New("syntax error")
	ErrNotFound          = errors.New("target struct is not found")
	ErrDone              = errors.New("Done")
	ErrInvalidConfigFile = errors.New(ErrMsgInvalidConfigFile)

	ImportOptions = &imports.Options{
		FormatOnly: false,
		TabIndent:  true,
		Comments:   true,
	}
)
