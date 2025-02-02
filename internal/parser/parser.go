package parser

import (
	"bufio"
	"errors"
	"go/parser"
	"go/token"
	"os"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
)

type parserFn func(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error

// ParseBuilderFile creates a file based on config and returns the first encountered error.
func ParseBuilderFile(config *cmd.Config) (*generation.StructGenHelper, error) {
	if config.WithValidation {
		fset := token.NewFileSet()
		if _, err := parser.ParseFile(fset, config.Source, nil, 0); err != nil {
			return nil, err
		}
	}

	file, err := os.Open(config.Source)
	if err != nil {
		return nil, err
	}

	structHelper := &generation.StructGenHelper{
		Package: config.Package,
		Name:    config.Name,
	}
	scanner := bufio.NewReader(file)

	parserFn := getParserMode(config.ParserMode)
	if parserFn == nil {
		return nil, consts.ErrInvalidParserMode
	}

	if err := parserFn(config, scanner, structHelper); err != nil && !errors.Is(err, consts.ErrDone) {
		return nil, err
	}

	if len(structHelper.Name) == 0 {
		return nil, consts.ErrNoStructsFound
	}

	if len(structHelper.Package) == 0 {
		return nil, consts.ErrPackageNotFound
	}

	return structHelper, nil
}

func getParserMode(parserMode consts.Mode) parserFn {
	switch parserMode {
	case consts.MODE_AST:
		return parseDataByAST
	case consts.MODE_FAST:
		return parseDataByCustomParser
	default:
		return nil
	}
}
