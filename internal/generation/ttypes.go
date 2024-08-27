package generation

import (
	"fmt"
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
	Name        string
	Package     string
	Fields      []*Field
	Imports     []string
	maxFieldLen int
}

func (s *StructGenHelper) ToSource() string {
	if s.maxFieldLen == 0 {
		for _, f := range s.Fields {
			if len(f.Name) > s.maxFieldLen {
				s.maxFieldLen = len(f.Name)
			}
		}
		fmt.Println(s.maxFieldLen)
	}
	srcBuilder := strings.Builder{}
	srcBuilder.WriteString(consts.BUILD_HEADER)
	srcBuilder.WriteString("\n")
	srcBuilder.WriteString(consts.BUILD_PACKAGE)
	srcBuilder.WriteString(" ")
	srcBuilder.WriteString(s.Package)

	srcBuilder.WriteString("\n\nimport (\n")
	for _, importVal := range s.Imports {
		srcBuilder.WriteString("\t")
		srcBuilder.WriteString(importVal)
		srcBuilder.WriteString("\n")
	}

	srcBuilder.WriteString(")\n\ntype ")
	srcBuilder.WriteString(s.Name)
	srcBuilder.WriteString("Builder struct {\n")
	for _, field := range s.Fields {
		srcBuilder.WriteString("\t")
		srcBuilder.WriteString(field.Name)
		srcBuilder.WriteString(strings.Repeat(" ", s.maxFieldLen-len(field.Name)+1))
		srcBuilder.WriteString(field.Type)
		if len(field.Tags) > 0 {
			srcBuilder.WriteString(" ")
			srcBuilder.WriteString(field.Tags)
		}
		srcBuilder.WriteString("\n")
	}
	srcBuilder.WriteString("}\n\n")

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
	builder.WriteString("Builder {\n\tif b == nil {\n\t\treturn nil\n\t}\n\n\treturn &")
	builder.WriteString(s.Name)
	builder.WriteString("Builder{")
	for _, field := range s.Fields {
		builder.WriteString("\n\t\t")
		builder.WriteString(field.Name)
		builder.WriteString(":")
		builder.WriteString(strings.Repeat(" ", s.maxFieldLen-len(field.Name)+1))
		builder.WriteString("b.")
		builder.WriteString(field.Name)
		builder.WriteString(",")
	}
	builder.WriteString("\n\t}\n}\n\n")

	return builder.String()
}

func (s *StructGenHelper) genMethod(field *Field) string {
	paramName := strings.ToLower(field.Name)
	if utils.IsKeyword(paramName) {
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
	builder.WriteString("\n\treturn b\n}\n\n")

	return builder.String()
}

func (s *StructGenHelper) genBuildMethod() string {
	builder := strings.Builder{}
	builder.WriteString("func (b *")
	builder.WriteString(s.Name)
	builder.WriteString("Builder) Build() *")
	builder.WriteString(s.Name)
	builder.WriteString(" {\n\treturn &")
	builder.WriteString(s.Name)
	builder.WriteString("{\n")

	for _, field := range s.Fields {
		builder.WriteString("\t\t")
		builder.WriteString(field.Name)
		builder.WriteString(":")

		builder.WriteString(strings.Repeat(" ", s.maxFieldLen-len(field.Name)+1))
		builder.WriteString("b.")
		builder.WriteString(field.Name)
		builder.WriteString(",\n")
	}

	builder.WriteString("\t}\n}\n")

	return builder.String()
}
