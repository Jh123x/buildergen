package generation

import (
	"fmt"
	"go/ast"
	"go/token"
	"log"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
)

// GenerateBuilder generates the builder source code based on the given arguments.
func GenerateBuilder(tSet *token.FileSet, typeSpec *ast.TypeSpec, imports []string, config *cmd.Config) (string, error) {
	structHelper := &StructGenHelper{
		Name:    config.Name,
		Package: config.Package,
		Fields:  make([]*Field, 0, 1000),
		Imports: imports,
	}

	if typeSpec.Type != nil {
		if typed, ok := typeSpec.Type.(*ast.StructType); ok {
			if err := generateStructFields(structHelper, typed); err != nil {
				return "", err
			}
		}
	}

	return structHelper.ToSource(), nil
}

func generateStructFields(helper *StructGenHelper, structs *ast.StructType) error {
	for _, field := range structs.Fields.List {
		name, err := getName(field.Names)
		if err != nil {
			return err
		}

		typeVal, err := getType(field.Type)
		if err != nil {
			return err
		}

		helper.Fields = append(helper.Fields, &Field{
			Name: name,
			Type: typeVal,
			Tags: getTag(field.Tag),
		})
	}

	return nil
}

func getTag(tag *ast.BasicLit) string {
	if tag == nil {
		return ""
	}
	return tag.Value
}

func getType(typeVal ast.Expr) (string, error) {
	switch v := typeVal.(type) {
	case *ast.Ident:
		return v.Name, nil
	case *ast.StarExpr:
		name, err := getType(v.X)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("*%s", name), nil
	case *ast.ArrayType:
		name, err := getType(v.Elt)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("[]%s", name), nil

	case *ast.MapType:
		keyType, err := getType(v.Key)
		if err != nil {
			return "", err
		}
		valType, err := getType(v.Value)
		if err != nil {
			return "", err
		}

		return fmt.Sprintf("map[%s]%s", keyType, valType), nil
	case *ast.SelectorExpr:
		pkg, err := getType(v.X)
		if err != nil {
			return "", nil
		}
		sType, err := getType(v.Sel)
		if err != nil {
			return "", nil
		}

		return fmt.Sprintf("%s.%s", pkg, sType), nil
	default:
		log.Println(v)
		return "", consts.ErrTypeNotfound
	}

}

func getName(idents []*ast.Ident) (string, error) {
	for _, val := range idents {
		if len(val.Name) > 0 {
			return val.Name, nil
		}
	}

	return "", consts.ErrNameNotFound
}
