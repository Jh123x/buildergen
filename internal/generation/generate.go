package generation

import (
	"go/ast"
	"log"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
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

		if err := getType(field.Type, &builder); err != nil {
			return err
		}

		helper.Fields = append(
			helper.Fields,
			&Field{
				Name: name,
				Type: builder.String(),
				Tags: getTag(field.Tag),
			},
		)
	}

	return nil
}

func getTag(tag *ast.BasicLit) string {
	if tag == nil {
		return consts.EMPTY_STR
	}

	return tag.Value
}

func getType(typeVal ast.Expr, builder *strings.Builder) error {
	switch v := typeVal.(type) {
	case *ast.Ident:
		builder.WriteString(v.Name)
		return nil
	case *ast.StarExpr:
		builder.WriteString("*")
		if err := getType(v.X, builder); err != nil {
			return err
		}
	case *ast.ArrayType:
		builder.WriteString("[]")
		if err := getType(v.Elt, builder); err != nil {
			return err
		}
	case *ast.MapType:
		builder.WriteString("map[")
		if err := getType(v.Key, builder); err != nil {
			return err
		}

		builder.WriteString("]")
		if err := getType(v.Value, builder); err != nil {
			return err
		}
	case *ast.SelectorExpr:
		if err := getType(v.X, builder); err != nil {
			return err
		}

		builder.WriteString(".")
		if err := getType(v.Sel, builder); err != nil {
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
