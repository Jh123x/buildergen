package consts

import "errors"

const (
	ErrMsgSrcNotfound    = "source file is required"
	ErrMsgNoStructsFound = "source file has no structs"
	ErrMsgNameNotFound   = "name is required"
	ErrMsgTypeNotFound   = "type not found for field"
	ErrMsgNotGoFile      = "source is not a valid go file"
	ErrMsgInvalidStruct  = "invalid struct type"
)

var (
	ErrSrcNotFound       = errors.New(ErrMsgSrcNotfound)
	ErrNoStructsFound    = errors.New(ErrMsgNoStructsFound)
	ErrNameNotFound      = errors.New(ErrMsgNameNotFound)
	ErrTypeNotfound      = errors.New(ErrMsgTypeNotFound)
	ErrNotGoFile         = errors.New(ErrMsgNotGoFile)
	ErrInvalidStructType = errors.New(ErrMsgInvalidStruct)
)
