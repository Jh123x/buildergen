package writer

import (
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
)

func MultiFileWrite(path string, structs ...*generation.StructGenHelper) error {
	switch len(structs) {
	case 0:
		return nil
	case 1:
		return WriteToSingleFile(path, structs[0].ToSource())
	default:
		break
	}

	structs = utils.FilterNil(structs)

	pkgName, err := mergePackages(structs)
	if err != nil {
		return err
	}

	finalImports, err := mergeImports(structs)
	if err != nil {
		return err
	}

	writeHelper := writeHelper{
		pkg:     pkgName,
		imports: finalImports,
		structs: structs,
	}

	return WriteToSingleFile(path, writeHelper.ToSource())
}
