package cmd

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jh123x/buildergen/pkg/consts"
	"github.com/Jh123x/buildergen/pkg/generation"
)

var (
	ErrInvalidConfigFile = consts.ErrInvalidConfigFile
	ErrSrcNotFound       = consts.ErrSrcNotFound
	ErrNameNotFound      = consts.ErrNameNotFound
	ErrNotGoFile         = consts.ErrNotGoFile
)

type PrinterFn func(string, ...any) (int, error)

var _ PrinterFn = fmt.Printf

//go:generate buildergen -src=./ttypes.go -name=Config

type Config struct {
	Source         string      `yaml:"source"`
	Destination    string      `yaml:"destination"`
	Package        string      `yaml:"package"`
	Name           string      `yaml:"name"`
	WithValidation bool        `yaml:"with-validation"`
	ParserMode     consts.Mode `yaml:"mode"`
	IncGenCmd      bool        `yaml:"inc-cmd"`
	generationCmd  string
}

type ConfigChan struct {
	StructHelper *generation.StructGenHelper
	Destination  string
	Err          error
}

type Field struct {
	Name string
	Type string
	Tags string
}

type Import struct {
	Name string
	Path string
}

func (i *Import) ToImport() string {
	if len(i.Name) == 0 {
		return i.Path
	}

	if i.Name+"\"" == filepath.Base(i.Path) {
		return i.Path
	}

	return fmt.Sprintf("%s %s", i.Name, i.Path)
}

func (i *Import) GetName() string {
	if len(i.Name) == 0 {
		return filepath.Base(i.Path[1 : len(i.Path)-1])
	}

	return i.Name
}

// NewConfig creates a new config with the given arguments.
// It also initializes the default values config arguments.
func NewConfig(src, dst, pkg, name string, validation, incCmd bool, parserMode consts.Mode) (*Config, error) {
	config := &Config{
		Source:         src,
		Name:           name,
		Package:        pkg,
		Destination:    dst,
		WithValidation: validation,
		ParserMode:     parserMode,
		IncGenCmd:      incCmd,
	}

	generationCmd, err := config.ToCommand()
	if err != nil {
		return nil, err
	}

	if _, err := config.FillDefaults(); err != nil {
		return nil, err
	}

	config.generationCmd = generationCmd
	return config, nil
}

func (c *Config) FillDefaults() (*Config, error) {
	if c == nil {
		return nil, ErrInvalidConfigFile
	}

	if c.Source == "" {
		return nil, ErrSrcNotFound
	}

	if c.Name == "" {
		return nil, ErrNameNotFound
	}

	if !strings.HasSuffix(c.Source, ".go") {
		return nil, ErrNotGoFile
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

func (c *Config) ToCommand() (string, error) {
	if c == nil {
		return "", ErrInvalidConfigFile
	}

	if c.Source == "" {
		return "", ErrSrcNotFound
	}

	if c.Name == "" {
		return "", ErrNameNotFound
	}

	cmd := strings.Builder{}
	cmd.WriteString("buildergen")
	cmd.WriteString(" --src=")
	cmd.WriteString(c.Source)
	cmd.WriteString(" --name=")
	cmd.WriteString(c.Name)

	if c.Destination != "" {
		cmd.WriteString(" --dst=")
		cmd.WriteString(c.Destination)
	}

	if c.Package != "" {
		cmd.WriteString(" --pkg=")
		cmd.WriteString(c.Package)
	}

	if c.WithValidation {
		cmd.WriteString(" --validation")
	}

	return cmd.String(), nil
}
