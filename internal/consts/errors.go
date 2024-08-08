package consts

import "errors"

const (
	ErrMsgSrcNotfound    = "source file is required"
	ErrMsgNoStructsFound = "source file has no structs"
	ErrMsgNameNotFound   = "name is required"
	ErrMsgTypeNotFound   = "type not found for field"
)

var (
	ErrSrcNotFound    = errors.New(ErrMsgSrcNotfound)
	ErrNoStructsFound = errors.New(ErrMsgNoStructsFound)
	ErrNameNotFound   = errors.New(ErrMsgNameNotFound)
	ErrTypeNotfound   = errors.New(ErrMsgTypeNotFound)
)
