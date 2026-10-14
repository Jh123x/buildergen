package cmd

import (
	"fmt"

	"github.com/Jh123x/buildergen/pkg/common"
	"github.com/Jh123x/buildergen/pkg/model"
)

var (
	ErrInvalidConfigFile = common.ErrInvalidConfigFile
	ErrSrcNotFound       = common.ErrSrcNotFound
	ErrNameNotFound      = common.ErrNameNotFound
	ErrNotGoFile         = common.ErrNotGoFile
)

type PrinterFn func(string, ...any) (int, error)

var _ PrinterFn = fmt.Printf

//go:generate buildergen -src=./ttypes.go -name=Config

type Config = model.Config

type ConfigChan = model.ConfigChan

type Field = model.Field

type Import = model.Import

// NewConfig creates a new config with the given arguments.
// It also initializes the default values config arguments.
func NewConfig(src, dst, pkg, name string, validation, incCmd bool, parserMode model.Mode) (*model.Config, error) {
	config := &model.Config{
		Source:         src,
		Name:           name,
		Package:        pkg,
		Destination:    dst,
		WithValidation: validation,
		ParserMode:     parserMode,
		IncGenCmd:      incCmd,
	}

	if _, err := config.FillDefaults(); err != nil {
		return nil, err
	}

	generationCmd, err := config.ToCommand()
	if err != nil {
		return nil, err
	}

	config.GenerationCmd = generationCmd
	return config, nil
}
