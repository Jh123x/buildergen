package parser

import (
	"bufio"
	"go/parser"
	"go/token"
	"io"
	"os"
	"strings"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
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

	structHelper := &generation.StructGenHelper{}
	scanner := bufio.NewReader(file)

	parseData(config, scanner, structHelper)

	if len(structHelper.Name) == 0 {
		return consts.EMPTY_STR, consts.ErrNoStructsFound
	}

	if len(structHelper.Package) == 0 {
		return consts.EMPTY_STR, consts.ErrSyntaxErr
	}

	return structHelper.ToSource(), nil
}

func parseData(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	for {
		kw, err := scanner.ReadString(' ')
		if err == io.EOF {
			return consts.ErrNotFound
		}

		if err != nil {
			return err
		}

		kw = strings.Trim(kw, consts.DEFAULT_TRIM)
		if len(kw) == 0 {
			continue
		}

		if err := parseByKeyword(kw, scanner, helper, config); err != nil {
			return err
		}
	}
}

func parseByKeyword(kw string, scanner *bufio.Reader, helper *generation.StructGenHelper, config *cmd.Config) error {
	switch kw {
	case consts.KEYWORD_PACKAGE:
		if err := parsePkg(scanner, helper); err != nil {
			return err
		}
	case consts.KEYWORD_IMPORT:
		if err := parseImport(scanner, helper); err != nil {
			return err
		}
	case consts.KEYWORD_TYPE:
		if err := parseType(scanner, helper, config.Name); err != nil {
			if err == consts.ErrDone {
				return nil
			}

			return err
		}
	default:
		if strings.HasPrefix(kw, consts.COMMENTS) {
			if err := parseInlineComments(scanner); err != nil {
				return err
			}
		}

		if strings.HasSuffix(kw, consts.COMMENT_START) {
			if err := parseMultilineComments(scanner); err != nil {
				return err
			}
		}
	}

	return nil
}

func parseMultilineComments(scanner *bufio.Reader) error {
	for {
		if _, err := scanner.ReadString(consts.COMMENT_END[0]); err != nil {
			return err
		}

		v, err := scanner.ReadByte()
		if err != nil {
			return err
		}

		// End of comment
		if v == '/' {
			break
		}
	}

	return nil
}

func parseInlineComments(scanner *bufio.Reader) error {
	_, err := scanner.ReadString('\n')
	return err

}

func parseType(scanner *bufio.Reader, helper *generation.StructGenHelper, target string) error {
	name, err := scanner.ReadString(' ')
	if err != nil {
		return err
	}

	name = strings.Trim(name, consts.DEFAULT_TRIM)
	if name != target {
		return nil
	}

	helper.Name = name
	typeVal, err := scanner.ReadString(' ')

	if err != nil {
		return err
	}

	typeVal = strings.Trim(typeVal, consts.DEFAULT_TRIM)
	if typeVal != consts.KEYWORD_STRUCT {
		return nil
	}

	if err := parseStruct(scanner, helper); err != nil {
		return err
	}

	return consts.ErrDone
}

func parseStruct(scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	fields, err := scanner.ReadString('}')
	if err != nil {
		return err
	}

	fieldRows := strings.Split(fields, "\n")
	for _, row := range fieldRows {
		row = strings.Trim(row, consts.DEFAULT_TRIM)
		if len(row) <= 1 {
			continue
		}

		field, err := parseFieldRow(row)
		if err != nil {
			return err
		}

		helper.Fields = append(helper.Fields, field)
	}

	return nil
}

func parseFieldRow(row string) (*generation.Field, error) {
	tokens := utils.Filter(
		utils.Map(
			strings.SplitN(row, " ", 3),
			func(val string) string { return strings.Trim(val, consts.DEFAULT_TRIM) },
		),
		func(val string) bool { return len(val) > 0 },
	)

	switch len(tokens) {
	case 2:
		return &generation.Field{
			Name: tokens[0],
			Type: tokens[1],
		}, nil
	case 3:
		return &generation.Field{
			Name: tokens[0],
			Type: tokens[1],
			Tags: tokens[2],
		}, nil
	default:
		return nil, consts.ErrSyntaxErr
	}
}

func parsePkg(scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	pkgName, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}

	helper.Package = strings.Trim(pkgName, consts.DEFAULT_TRIM)
	return nil
}

func parseImport(scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	impVal, err := scanner.ReadString('\n')
	if err != nil {
		return err
	}
	impVal = strings.Trim(impVal, consts.DEFAULT_TRIM)

	if !strings.Contains(impVal, "(") {
		imp, err := parseImportLine(impVal)
		if err != nil {
			return err
		}

		helper.Imports = []*generation.Import{imp}
		return nil
	}

	importLines, err := scanner.ReadString(')')
	if err != nil {
		return err
	}

	importRows := strings.Split(importLines, "\n")
	helper.Imports = make([]*generation.Import, 0, len(importRows[:len(importRows)-1]))

	for _, row := range importRows {
		if len(row) <= 2 {
			continue
		}

		imp, err := parseImportLine(strings.Trim(row, consts.DEFAULT_TRIM))
		if err != nil {
			return err
		}
		helper.Imports = append(helper.Imports, imp)
	}

	return nil
}

func parseImportLine(importLine string) (*generation.Import, error) {
	tokens := strings.Split(importLine, " ")
	switch len(tokens) {
	case 0:
		return nil, nil
	case 1:
		return &generation.Import{Path: tokens[0]}, nil
	case 2:
		return &generation.Import{
			Name: tokens[0],
			Path: tokens[1],
		}, nil
	default:
		return nil, consts.ErrSyntaxErr
	}
}
