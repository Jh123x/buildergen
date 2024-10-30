package parser

import (
	"bufio"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
)

var _ parserFn = parseDataByAST

func parseDataByAST(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, config.Source, nil, 0)
	if err != nil {
		return err
	}

	if len(helper.Package) == 0 && astFile.Package.IsValid() {
		helper.Package = astFile.Name.Name
	}

	res, ok := findRequestedStructType(astFile, config.Name)
	if !ok {
		return consts.ErrNoStructsFound
	}

	helper.Imports = parseData(astFile.Imports)

	if err := generation.GenerateBuilder(res, helper); err != nil {
		return err
	}

	return nil
}

func parseData(imports []*ast.ImportSpec) []*generation.Import {
	result := make([]*generation.Import, 0, len(imports))

	for _, res := range imports {
		if res.Name == nil {
			result = append(result, &generation.Import{Path: res.Path.Value})
			continue
		}

		result = append(result, &generation.Import{
			Name: res.Name.String(),
			Path: res.Path.Value,
		})
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
