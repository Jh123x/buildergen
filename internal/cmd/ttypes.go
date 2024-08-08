package cmd

import (
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
)

type Config struct {
	Source      string
	Destination string
	Package     string
	Name        string
}

func NewConfig(src, dst, pkg, name *string) (*Config, error) {
	if utils.IsNilOrEmpty(src) {
		return nil, consts.ErrSrcNotFound
	}

	if utils.IsNilOrEmpty(name) {
		return nil, consts.ErrNameNotFound
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
