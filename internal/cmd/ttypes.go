package cmd

import (
	"fmt"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
)

type PrinterFn func(string, ...any) (int, error)

var _ PrinterFn = fmt.Printf

//go:generate buildergen -src=./ttypes.go -name Config

type Config struct {
	Source      string
	Destination string
	Package     string
	Name        string
}

// NewConfig creates a new config with the given arguments.
// It also initializes the default values config arguments.
func NewConfig(src, dst, pkg, name *string) (*Config, error) {
	if utils.IsNilOrEmpty(src) {
		return nil, consts.ErrSrcNotFound
	}

	if utils.IsNilOrEmpty(name) {
		return nil, consts.ErrNameNotFound
	}

	if !strings.HasSuffix(*src, ".go") {
		return nil, consts.ErrNotGoFile
	}

	if utils.IsNilOrEmpty(dst) {
		srcFileName := strings.Split(*src, ".")
		fileName := strings.Join(srcFileName[:len(srcFileName)-1], ".")
		builder := strings.Builder{}
		builder.WriteString(fileName)
		builder.WriteString(consts.DEFAULT_BUILDER_SUFFIX)
		derivedDst := builder.String()
		dst = &derivedDst
	}

	if utils.IsNilOrEmpty(pkg) {
		emptyStr := ""
		pkg = &emptyStr
	}

	return &Config{
		Source:      *src,
		Destination: *dst,
		Package:     *pkg,
		Name:        *name,
	}, nil
}
