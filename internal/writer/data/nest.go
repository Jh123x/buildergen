package data

import (
	"context"
	"os"
)

type Test struct {
	Val          string
	ImportedType *os.FileMode
}

type OtherStruct struct {
	OtherImports context.Context
}
