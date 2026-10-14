package generation

import (
	"go/ast"
	"log"
	"strings"

	"github.com/Jh123x/buildergen/pkg/common"
	"github.com/Jh123x/buildergen/pkg/model"
)

// GenerateBuilder generates the builder source code based on the given arguments.
func GenerateBuilder(typeSpec *ast.TypeSpec, structHelper *model.StructGenHelper) error {
	if typeSpec.Type == nil {
		return common.ErrNoStructsFound
	}

	typed, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return common.ErrInvalidStructType
	}

	if err := generateStructFields(structHelper, typed); err != nil {
		return err
	}

	return nil
}

func generateStructFields(helper *model.StructGenHelper, structs *ast.StructType) error {
	for _, field := range structs.Fields.List {
		if len(field.Names) == 0 {
			continue
		}

		name := field.Names[0].Name
		var fieldRes model.Field

		if err := getType(field.Type, &strings.Builder{}, &fieldRes, false, ""); err != nil {
			return err
		}

		fieldRes.Name = name

		if err := getName(field.Names, &strings.Builder{}); err != nil {
			return err
		}

		helper.Fields = append(helper.Fields, &fieldRes)
	}

	return nil
}

func getType(expr ast.Expr, builder *strings.Builder, fieldRes *model.Field, isMoved bool, basePkg string) error {
	switch t := expr.(type) {
	case *ast.Ident:
		v := t.Name
		builder.WriteString(v)
		fieldRes.Type = v
	case *ast.SelectorExpr:
		x, ok := t.X.(*ast.Ident)
		if !ok {
			return common.ErrInvalidStructType
		}
		v := x.Name + "." + t.Sel.Name
		builder.WriteString(v)
		fieldRes.Type = v
	case *ast.StarExpr:
		err := getType(t.X, builder, fieldRes, isMoved, basePkg)
		if err != nil {
			return err
		}
		builder.WriteString("*")
		fieldRes.Type = "*" + fieldRes.Type
	default:
		log.Println("unsupported type")
		return common.ErrTypeNotfound
	}

	return nil
}

func getName(idents []*ast.Ident, builder *strings.Builder) error {
	if len(idents) == 0 {
		return common.ErrNameNotFound
	}
	builder.WriteString(idents[0].Name)
	return nil
}
