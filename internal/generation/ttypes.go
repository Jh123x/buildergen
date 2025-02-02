package generation

import (
	"fmt"
	"path"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/utils"
)

type Import struct {
	Name string
	Path string
}

func (i *Import) ToImport() string {
	if len(i.Name) == 0 {
		return i.Path
	}

	return fmt.Sprintf("%s %s", i.Name, i.Path)
}

func (i *Import) GetName() string {
	if len(i.Name) == 0 {
		return path.Base(i.Path[1 : len(i.Path)-1])
	}

	return i.Name
}

type Field struct {
	Name string
	Type string
	Tags string
}

func (f *Field) GetUsedPackageName() string {
	if !strings.Contains(f.Type, ".") {
		return consts.EMPTY_STR
	}

	name := strings.SplitN(f.Type, ".", 2)[0]
	return strings.TrimPrefix(name, "*")
}

type StructGenHelper struct {
	Name    string
	Package string
	Fields  []*Field
	Imports []*Import

	// Used Internally
	maxFieldLen  int
	maxTypeLen   int
	usedPackages map[string]consts.Empty
}

func (s *StructGenHelper) ToSource() string {
	if len(s.usedPackages) == 0 {
		s.usedPackages = make(map[string]consts.Empty, len(s.Fields))

		for _, f := range s.Fields {
			pkgName := f.GetUsedPackageName()
			if len(pkgName) == 0 {
				continue
			}

			s.usedPackages[pkgName] = consts.Empty{}
		}
	}

	srcBuilder := strings.Builder{}
	srcBuilder.WriteString(consts.BUILD_HEADER)
	srcBuilder.WriteString("\n")
	srcBuilder.WriteString(consts.BUILD_PACKAGE)
	srcBuilder.WriteString(" ")
	srcBuilder.WriteString(s.Package)

	if len(s.usedPackages) > 0 {
		importBuffer := make([]string, 0, len(s.Imports))
		for _, importVal := range s.Imports {
			if _, ok := s.usedPackages[importVal.GetName()]; !ok {
				continue
			}

			importBuffer = append(importBuffer, importVal.ToImport())
		}

		if len(importBuffer) == 1 {
			srcBuilder.WriteString("\n\nimport ")
			srcBuilder.WriteString(importBuffer[0])
		}

		if len(importBuffer) > 1 {
			srcBuilder.WriteString("\n\nimport (\n")
			for _, val := range s.Imports {
				srcBuilder.WriteString("\t")
				srcBuilder.WriteString(val.ToImport())
				srcBuilder.WriteString("\n")
			}
			srcBuilder.WriteString(")")
		}
	}

	srcBuilder.WriteString("\n\n")
	srcBuilder.WriteString(s.BuildStruct())

	return srcBuilder.String()
}

func (s *StructGenHelper) BuildStruct() string {
	if s.maxFieldLen == 0 {
		for _, f := range s.Fields {
			if len(f.Name) > s.maxFieldLen {
				s.maxFieldLen = len(f.Name)
			}

			if len(f.Type) > s.maxTypeLen {
				s.maxTypeLen = len(f.Type)
			}
		}
	}

	srcBuilder := strings.Builder{}
	srcBuilder.WriteString("type ")
	srcBuilder.WriteString(s.Name)
	srcBuilder.WriteString("Builder struct {\n")
	for _, field := range s.Fields {
		srcBuilder.WriteString("\t")
		srcBuilder.WriteString(field.Name)
		srcBuilder.WriteString(strings.Repeat(" ", s.maxFieldLen-len(field.Name)+1))
		srcBuilder.WriteString(field.Type)
		if len(field.Tags) > 0 {
			srcBuilder.WriteString(strings.Repeat(" ", s.maxTypeLen-len(field.Type)+1))
			srcBuilder.WriteString(field.Tags)
		}
		srcBuilder.WriteString("\n")
	}
	srcBuilder.WriteString("}\n\n")

	s.genNewMethod(&srcBuilder)

	for _, field := range s.Fields {
		s.genMethod(&srcBuilder, field)
	}

	s.genBuildMethod(&srcBuilder)

	return srcBuilder.String()
}

func (s *StructGenHelper) genNewMethod(builder *strings.Builder) {
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
}

func (s *StructGenHelper) genMethod(builder *strings.Builder, field *Field) string {
	paramName := utils.LowerFirstLetter(field.Name)
	if utils.IsKeyword(paramName) {
		paramName += "_"
	}

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

func (s *StructGenHelper) genBuildMethod(builder *strings.Builder) string {
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
