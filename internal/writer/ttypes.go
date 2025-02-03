package writer

import (
	"sort"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
)

type importHelper struct {
	importName  string
	structNames utils.Set[string]
}

type writeHelper struct {
	pkg     string
	imports []*generation.Import

	// Imports and Package filed will be ignored.
	structs []*generation.StructGenHelper

	// Used Internally
	usedPackages utils.Set[string]
}

func (w *writeHelper) ToSource() string {
	builder := strings.Builder{}

	builder.WriteString(consts.BUILD_HEADER)
	builder.WriteString("\n")
	builder.WriteString("package ")
	builder.WriteString(w.pkg)

	switch len(w.imports) {
	case 1:
		builder.WriteString("\n\nimport ")
		builder.WriteString(w.imports[0].ToImport())
		builder.WriteString("\n\n")
	case 0:
		break
	default:
		builder.WriteString("\n\nimport ")
		builder.WriteString("(\n")

		importAcc := utils.Map(w.imports, func(i *generation.Import) string { return i.ToImport() })
		sort.Strings(importAcc)

		for _, importStr := range importAcc {
			builder.WriteString("\t")
			builder.WriteString(importStr)
			builder.WriteString("\n")
		}
		builder.WriteString(")")
		builder.WriteString("\n\n")
	}

	for idx, s := range w.structs {
		builder.WriteString(s.BuildStruct())
		if idx < len(w.structs)-1 {
			builder.WriteString("\n")
		}
	}

	return builder.String()
}
