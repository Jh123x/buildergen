package parser

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/Jh123x/buildergen/pkg/cmd"
	"github.com/Jh123x/buildergen/pkg/consts"
	"github.com/Jh123x/buildergen/pkg/generation"
)

var _ parserFn = parseDataByAST

func parseDataByAST(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	file, err := parseGoFile(config.Source)
	if err != nil {
		return err
	}

	if !file.Package.IsValid() {
		return fmt.Errorf("invalid package name")
	}

	if !config.IncGenCmd {
		helper.GenerationCmd = ""
	}
	helper.SrcPackage = file.Name.Name
	helper.DstPackage = helper.SrcPackage
	if config.Package != "" {
		helper.DstPackage = config.Package
	}

	res, ok := findRequestedStructType(file, config.Name)
	if !ok {
		return consts.ErrNoStructsFound
	}

	importFiles := parseData(file.Imports)
	if helper.SrcPackage != helper.DstPackage {
		importPath, err := importPathFromFile(config.Source)
		if err != nil {
			return err
		}

		importFiles = append(importFiles, &generation.Import{
			Name: helper.SrcPackage,
			Path: "\"" + importPath + "\"",
		})
	}

	helper.Imports = importFiles
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

func parseGoFile(path string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, path, nil, 0)
}
