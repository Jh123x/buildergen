package writer

import (
	"fmt"
	"path"
	"sort"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
)

func mergeImports(structs []*generation.StructGenHelper) ([]*generation.Import, error) {
	structs = utils.FilterNil(structs)
	imports := make(map[string]importHelper, 1000)

	for _, s := range structs {
		if s == nil {
			continue
		}

		usedPackages := s.GetUsedPackages()
		for _, i := range s.Imports {
			path := i.Path
			prevImport, ok := imports[path]
			if !usedPackages.Has(i.GetName()) {
				continue
			}
			if ok && prevImport.importName != i.GetName() {
				structsNames := prevImport.structNames.ToList()
				sort.Strings(structsNames)

				return nil, fmt.Errorf(
					"import names for package='%s' are inconsistent between %s (%s) and %s (%s)",
					path,
					strings.Join(structsNames, ", "),
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
	for importPath, h := range imports {
		imp := &generation.Import{Path: importPath}
		if path.Base(importPath[1:len(importPath)-1]) != h.importName {
			imp.Name = h.importName
		}
		acc = append(acc, imp)
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

	pkgSet := make(utils.Set[string], len(structs))
	for _, s := range structs {
		if s == nil || s.Package == "" {
			continue
		}

		pkgSet.Add(s.Package)
	}

	switch len(pkgSet) {
	case 0:
		structNames := utils.Map(
			structs,
			func(s *generation.StructGenHelper) string { return s.Name },
		)
		sort.Strings(structNames)
		return consts.EMPTY_STR, fmt.Errorf(
			"no packages found within structs: %s",
			strings.Join(structNames, ", "),
		)
	case 1:
		return structs[0].Package, nil
	default:
		break
	}

	data := pkgSet.ToList()
	sort.Strings(data)

	return consts.EMPTY_STR, fmt.Errorf(
		"multiple packages found within the same file: %s",
		strings.Join(data, ", "),
	)
}
