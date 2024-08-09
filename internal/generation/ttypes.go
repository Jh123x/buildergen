package generation

import (
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
)

type Field struct {
	Name string
	Type string
	Tags string
}

type StructGenHelper struct {
	Name    string
	Package string
	Fields  []*Field
	Imports []string
}

func (s *StructGenHelper) ToSource() string {
	srcBuilder := strings.Builder{}
	srcBuilder.WriteString(consts.BUILD_HEADER)
	srcBuilder.WriteString("\n")
	srcBuilder.WriteString(consts.BUILD_PACKAGE)
	srcBuilder.WriteString(" ")
	srcBuilder.WriteString(s.Package)

	srcBuilder.WriteString("\n\nimport (\n")
	for _, importVal := range s.Imports {
		srcBuilder.WriteString(importVal)
		srcBuilder.WriteString("\n")
	}

	srcBuilder.WriteString(")\n\ntype ")
	srcBuilder.WriteString(s.Name)
	srcBuilder.WriteString("Builder struct {\n")
	for _, field := range s.Fields {
		srcBuilder.WriteString("\t")
		srcBuilder.WriteString(field.Name)
		srcBuilder.WriteString(" ")
		srcBuilder.WriteString(field.Type)
		if len(field.Tags) > 0 {
			srcBuilder.WriteString(" ")
			srcBuilder.WriteString(field.Tags)
		}
		srcBuilder.WriteString("\n")
	}
	srcBuilder.WriteString("}\n")

	srcBuilder.WriteString(s.genNewMethod())

	for _, field := range s.Fields {
		srcBuilder.WriteString(s.genMethod(field))
	}

	srcBuilder.WriteString(s.genBuildMethod())

	return srcBuilder.String()
}

func (s *StructGenHelper) genNewMethod() string {
	builder := strings.Builder{}
	builder.WriteString("func New")
	builder.WriteString(s.Name)
	builder.WriteString("Builder(b *")
	builder.WriteString(s.Name)
	builder.WriteString(") *")
	builder.WriteString(s.Name)
	builder.WriteString("Builder {\nif b == nil {\nreturn nil\n}\n\nreturn &")
	builder.WriteString(s.Name)
	builder.WriteString("Builder{\n")
	for _, field := range s.Fields {
		builder.WriteString(field.Name)
		builder.WriteString(": b.")
		builder.WriteString(field.Name)
		builder.WriteString(",\n")
	}
	builder.WriteString("\n}\n}\n\n")

	return builder.String()
}

func (s *StructGenHelper) genMethod(field *Field) string {
	paramName := strings.ToLower(field.Name)
	if utils.Contains(consts.Keywords, paramName) {
		paramName += "_"
	}

	builder := strings.Builder{}
	builder.WriteString("func (b *")
	builder.WriteString(s.Name)
	builder.WriteString("Builder) With")
	builder.WriteString(field.Name)
	builder.WriteString("(")
	builder.WriteString(paramName)
	builder.WriteString(" ")
	builder.WriteString(field.Type)
	builder.WriteString(") *")
	builder.WriteString(s.Name)
	builder.WriteString("Builder {\n\tb.")
	builder.WriteString(field.Name)
	builder.WriteString(" = ")
	builder.WriteString(paramName)
	builder.WriteString("\nreturn b\n}\n\n")

	return builder.String()
}

func (s *StructGenHelper) genBuildMethod() string {
	builder := strings.Builder{}
	builder.WriteString("func (b *")
	builder.WriteString(s.Name)
	builder.WriteString("Builder) Build() *")
	builder.WriteString(s.Name)
	builder.WriteString(" {")
	builder.WriteString("return &")
	builder.WriteString(s.Name)
	builder.WriteString("{\n")

	for _, field := range s.Fields {
		builder.WriteString(field.Name)
		builder.WriteString(": b.")
		builder.WriteString(field.Name)
		builder.WriteString(",\n")
	}

	builder.WriteString("}\n}\n\n")

	return builder.String()
}
