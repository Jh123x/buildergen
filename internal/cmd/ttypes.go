package cmd

import (
	"fmt"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
)

type PrinterFn func(string, ...any) (int, error)

var _ PrinterFn = fmt.Printf

//go:generate buildergen -src=./ttypes.go -name Config

type Config struct {
	Source         string      `yaml:"source"`
	Destination    string      `yaml:"destination"`
	Package        string      `yaml:"package"`
	Name           string      `yaml:"name"`
	WithValidation bool        `yaml:"with-validation"`
	ParserMode     consts.Mode `yaml:"mode"`
}

type ConfigChan struct {
	StructHelper *generation.StructGenHelper
	Destination  string
	Err          error
}

// NewConfig creates a new config with the given arguments.
// It also initializes the default values config arguments.
func NewConfig(src, dst, pkg, name string, validation bool, parserMode consts.Mode) (*Config, error) {
	config := &Config{
		Source:         src,
		Name:           name,
		Package:        pkg,
		Destination:    dst,
		WithValidation: validation,
		ParserMode:     parserMode,
	}

	return config.FillDefaults()
}

func (c *Config) FillDefaults() (*Config, error) {
	if c == nil {
		return nil, consts.ErrInvalidConfigFile
	}

	if c.Source == "" {
		return nil, consts.ErrSrcNotFound
	}

	if c.Name == "" {
		return nil, consts.ErrNameNotFound
	}

	if !strings.HasSuffix(c.Source, ".go") {
		return nil, consts.ErrNotGoFile
	}

	// Current default is empty string. Skip Package
	if c.Destination == "" {
		c.Destination = c.Source[:strings.LastIndex(c.Source, ".")] + consts.DEFAULT_BUILDER_SUFFIX
	}

	if c.ParserMode == "" {
		c.ParserMode = consts.MODE_AST
	}

	return c, nil
}
