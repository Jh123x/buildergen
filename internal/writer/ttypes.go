package writer

import (
	"strings"

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
}

func (w *writeHelper) ToSource() string {
	builder := strings.Builder{}

	builder.WriteString("package")
	builder.WriteString(w.pkg)
	builder.WriteString("\n\nimport(\n")

	for _, i := range w.imports {
		builder.WriteString("\t")
		builder.WriteString(i.ToImport())
		builder.WriteString("\n")
	}
	builder.WriteString(")\n\n")

	for _, s := range w.structs {
		builder.WriteString(s.BuildStruct())
	}

	return builder.String()
}
