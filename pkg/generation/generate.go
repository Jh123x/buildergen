package generation

import (
	"go/ast"
	"log"
	"slices"
	"strings"

	"github.com/Jh123x/buildergen/pkg/consts"
)

// GenerateBuilder generates the builder source code based on the given arguments.
func GenerateBuilder(typeSpec *ast.TypeSpec, structHelper *StructGenHelper) error {
	if typeSpec.Type == nil {
		return consts.ErrNoStructsFound
	}

	typed, ok := typeSpec.Type.(*ast.StructType)
	if !ok {
		return consts.ErrInvalidStructType
	}

	if err := generateStructFields(structHelper, typed); err != nil {
		return err
	}

	return nil
}

func generateStructFields(helper *StructGenHelper, structs *ast.StructType) error {
	for _, field := range structs.Fields.List {
		builder := strings.Builder{}
		if err := getName(field.Names, &builder); err != nil {
			return err
		}

		name := builder.String()
		builder.Reset()
		fieldRes := &Field{
			Name: name,
			Tags: getTag(field.Tag),
		}

		if err := getType(
			field.Type,
			&builder,
			fieldRes,
			helper.SrcPackage != helper.DstPackage,
			helper.SrcPackage,
		); err != nil {
			return err
		}
		fieldRes.Type = builder.String()
		helper.Fields = append(helper.Fields, fieldRes)
	}

	return nil
}

func getTag(tag *ast.BasicLit) string {
	if tag == nil {
		return consts.EMPTY_STR
	}

	return tag.Value
}

func getType(typeVal ast.Expr, builder *strings.Builder, fieldRes *Field, isMoved bool, basePkg string) error {
	switch v := typeVal.(type) {
	case *ast.Ident:
		if isMoved && !slices.Contains(consts.PrimitiveTypes, v.Name) && !strings.Contains(v.Name, ".") {
			builder.WriteString(basePkg)
			builder.WriteRune('.')
		}
		builder.WriteString(v.Name)
		return nil
	case *ast.StarExpr:
		builder.WriteString("*")
		if err := getType(v.X, builder, fieldRes, isMoved, basePkg); err != nil {
			return err
		}
	case *ast.ArrayType:
		builder.WriteString("[]")
		if err := getType(v.Elt, builder, fieldRes, isMoved, basePkg); err != nil {
			return err
		}
	case *ast.MapType:
		builder.WriteString("map[")
		if err := getType(v.Key, builder, fieldRes, isMoved, basePkg); err != nil {
			return err
		}

		builder.WriteString("]")
		if err := getType(v.Value, builder, fieldRes, isMoved, basePkg); err != nil {
			return err
		}
	case *ast.SelectorExpr:
		if err := getType(v.X, builder, fieldRes, false, basePkg); err != nil {
			return err
		}

		builder.WriteString(".")
		if err := getType(v.Sel, builder, fieldRes, false, basePkg); err != nil {
			return err
		}
	default:
		log.Println(v)
		return consts.ErrTypeNotfound
	}

	return nil
}

func getName(idents []*ast.Ident, builder *strings.Builder) error {
	for _, val := range idents {
		if len(val.Name) == 0 {
			continue
		}

		builder.WriteString(string(val.Name))
		return nil
	}

	return consts.ErrNameNotFound
}
