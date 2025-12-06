package parser

import (
	"bufio"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"golang.org/x/tools/go/packages"
)

var _ parserFn = parseDataByAST

func parseDataByAST(config *cmd.Config, scanner *bufio.Reader, helper *generation.StructGenHelper) error {
	fset := token.NewFileSet()
	astFile, err := parser.ParseFile(fset, config.Source, nil, 0)
	if err != nil {
		return err
	}

	if !astFile.Package.IsValid() {
		return fmt.Errorf("invalid package name")
	}

	helper.SrcPackage = astFile.Name.Name
	if config.Package != "" {
		helper.DstPackage = config.Package
	}

	res, ok := findRequestedStructType(astFile, config.Name)
	if !ok {
		return consts.ErrNoStructsFound
	}

	importFiles := parseData(astFile.Imports)
	if helper.SrcPackage != config.Package {
		res, err := importPathFromFile(config.Source)
		if err != nil {
			return err
		}

		importFiles = append(importFiles, &generation.Import{
			Name: helper.SrcPackage,
			Path: "\"" + res + "\"",
		})
	}

	helper.Imports = importFiles
	if err := generation.GenerateBuilder(res, helper); err != nil {
		return err
	}

	return nil
}

func importPathFromFile(filePath string) (string, error) {
	abs, err := filepath.Abs(filePath)
	if err != nil {
		return "", err
	}

	cfg := &packages.Config{
		Mode: packages.NeedFiles | packages.NeedName | packages.NeedModule,
	}

	// Load the package that owns this file.
	pkgs, err := packages.Load(cfg, "file="+abs)
	if err != nil {
		return "", err
	}

	if len(pkgs) == 0 {
		return "", fmt.Errorf("no package found for %s", filePath)
	}

	return pkgs[0].PkgPath, nil
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
