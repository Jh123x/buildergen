package parser

import (
	"bufio"
	"go/parser"
	"go/token"
	"os"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
)

type parserFn func(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error

var (
	parserMode = map[consts.Mode]parserFn{
		consts.MODE_FAST: parseDataByCustomParser,
		consts.MODE_AST:  parseDataByAST,
	}
)

// ParseBuilderFile creates a file based on config and returns the first encountered error.
func ParseBuilderFile(config *cmd.Config) (string, error) {
	if config.WithValidation {
		fset := token.NewFileSet()
		if _, err := parser.ParseFile(fset, config.Source, nil, 0); err != nil {
			return consts.EMPTY_STR, err
		}
	}

	file, err := os.Open(config.Source)
	if err != nil {
		return consts.EMPTY_STR, err
	}

	structHelper := &generation.StructGenHelper{
		Package: config.Package,
	}
	scanner := bufio.NewReader(file)

	parserFn, ok := parserMode[config.ParserMode]
	if !ok {
		return "", consts.ErrInvalidParserMode
	}

	if err := parserFn(config, scanner, structHelper); err != nil && err != consts.ErrDone {
		return "", err
	}

	if len(structHelper.Name) == 0 {
		return consts.EMPTY_STR, consts.ErrNoStructsFound
	}

	if len(structHelper.Package) == 0 {
		return consts.EMPTY_STR, consts.ErrSyntaxErr
	}

	return structHelper.ToSource(), nil
}
