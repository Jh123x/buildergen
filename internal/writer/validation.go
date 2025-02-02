package writer

import (
	"fmt"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
)

func mergeImports(structs []*generation.StructGenHelper) ([]*generation.Import, error) {
	imports := make(map[string]importHelper, 1000)

	for _, s := range structs {
		if s == nil {
			continue
		}
		for _, i := range s.Imports {
			path := i.Path
			prevImport, ok := imports[path]
			if ok && prevImport.importName != i.GetName() {
				return nil, fmt.Errorf(
					"import names for package='%s' are inconsistent between %s (%s) and %s (%s)",
					path,
					strings.Join(prevImport.structNames.ToList(), ", "),
					prevImport.importName,
					s.Name,
					i.GetName(),
				)
			}

			if !ok {
				imports[path] = importHelper{
					importName:  i.GetName(),
					structNames: utils.NewSet[string](),
				}
			}

			imports[path].structNames.Add(s.Name)
		}
	}

	acc := make([]*generation.Import, 0, len(imports))
	for path, h := range imports {
		acc = append(acc, &generation.Import{Path: path, Name: h.importName})
	}

	return acc, nil
}

func mergePackages(structs []*generation.StructGenHelper) (string, error) {
	if len(structs) == 0 {
		return consts.EMPTY_STR, nil
	}

	if len(structs) == 1 {
		return structs[0].Package, nil
	}

	pkgSet := make(map[string]consts.Empty, len(structs))
	for _, s := range structs {
		if s == nil {
			continue
		}

		pkgSet[s.Package] = consts.Empty{}
	}

	switch len(pkgSet) {
	case 0:
		return consts.EMPTY_STR, fmt.Errorf(
			"no packages found within structs: %s",
			strings.Join(
				utils.Map(
					structs,
					func(s *generation.StructGenHelper) string { return s.Name }),
				", ",
			),
		)
	case 1:
		return structs[0].Package, nil
	default:
		break
	}

	data := make([]string, 0, len(pkgSet))
	for n := range pkgSet {
		data = append(data, n)
	}

	return consts.EMPTY_STR, fmt.Errorf(
		"multiple packages found within the same file: %s",
		strings.Join(data, ", "),
	)
}
