package generation

import (
	"strings"

	"github.com/Jh123x/buildergen/pkg/consts"
	"github.com/Jh123x/buildergen/pkg/model"
)

type Field = model.Field

type StructGenHelper = model.StructGenHelper

func (s *StructGenHelper) ToSource() string {
	s.Preprocess()
	srcBuilder := strings.Builder{}
	srcBuilder.WriteString(consts.BUILD_HEADER)
	srcBuilder.WriteString("\n")
	srcBuilder.WriteString(consts.BUILD_PACKAGE)
	srcBuilder.WriteString(" ")
	srcBuilder.WriteString(s.DstPackage)

	usedPkgs := s.GetUsedPackages()
	if len(usedPkgs) > 0 {
		importBuffer := make([]string, 0, len(s.Imports))
		for _, importVal := range s.Imports {
			importName := importVal.GetName()
			if !usedPkgs.Has(importName) &&
				(s.DstPackage == s.SrcPackage || importName != s.SrcPackage) {
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
			for _, val := range importBuffer {
				srcBuilder.WriteString("\t")
				srcBuilder.WriteString(val)
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
	s.Preprocess()
	srcBuilder := strings.Builder{}
	if s.GenerationCmd != "" {
		srcBuilder.WriteString("// Generated using: ")
		srcBuilder.WriteString(s.GenerationCmd)
		srcBuilder.WriteString("\n")
	}
	srcBuilder.WriteString("type ")
	srcBuilder.WriteString(s.Name)
	srcBuilder.WriteString("Builder struct {\n")
	for _, field := range s.Fields {
		srcBuilder.WriteString("\t")
		srcBuilder.WriteString(field.Name)
		srcBuilder.WriteString(strings.Repeat(" ", s.MaxFieldLen-len(field.Name)+1))
		srcBuilder.WriteString(field.BuildType())
		if len(field.Tags) > 0 {
			srcBuilder.WriteString(strings.Repeat(" ", s.MaxTypeLen-len(field.Type)+1))
			srcBuilder.WriteString(field.Tags)
		}
		srcBuilder.WriteString("\n")
	}
	srcBuilder.WriteString("}\n\n")

	s.GenNewMethod(&srcBuilder)

	for _, field := range s.Fields {
		s.GenMethod(&srcBuilder, field)
	}

	s.GenBuildMethod(&srcBuilder)

	return srcBuilder.String()
}
