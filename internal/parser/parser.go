package parser

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"golang.org/x/tools/imports"
)

// ParseBuilderFile creates a file based on config and returns the first encountered error.
func ParseBuilderFile(config *cmd.Config) (string, error) {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, config.Source, nil, 0)
	if err != nil {
		return "", err
	}

	if len(config.Package) == 0 && astFile.Package.IsValid() {
		config.Package = astFile.Name.Name
	}

	res, ok := findRequestedStructType(astFile, config.Name)
	if !ok {
		return "", consts.ErrNoStructsFound
	}

	importData := parseData(astFile.Imports)
	results, err := generation.GenerateBuilder(fset, res, importData, config)
	if err != nil {
		return "", err
	}

	importRes, err := imports.Process("", []byte(results), consts.ImportOptions)
	if err != nil {
		return "", err
	}

	return string(importRes), nil
}

func parseData(imports []*ast.ImportSpec) []string {
	result := make([]string, 0, len(imports))

	for _, res := range imports {
		if res.Name == nil {
			result = append(result, res.Path.Value)
			continue
		}

		result = append(result, fmt.Sprintf("%s %s", res.Name, res.Path.Value))
	}

	return result
}

func findRequestedStructType(f *ast.File, structName string) (*ast.TypeSpec, bool) {
	for _, decl := range f.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || (genDecl.Tok != token.TYPE && genDecl.Tok != token.IMPORT) {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			if _, ok := typeSpec.Type.(*ast.StructType); ok && typeSpec.Name.Name == structName {
				return typeSpec, true
			}
		}
	}

	return nil, false
}
