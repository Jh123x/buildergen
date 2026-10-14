package model

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Jh123x/buildergen/pkg/utils"
)

// Mode represents the parser mode.
type Mode string

const (
	MODE_AST Mode = "DEFAULT"
)

var (
	ALL_MODES = []Mode{MODE_AST}
)

// Config represents the configuration for a single builder generation task.
type Config struct {
	Source         string `yaml:"source"`
	Destination    string `yaml:"destination"`
	Package        string `yaml:"package"`
	Name           string `yaml:"name"`
	WithValidation bool   `yaml:"with-validation"`
	ParserMode     Mode   `yaml:"mode"`
	IncGenCmd      bool   `yaml:"inc-cmd"`
	GenerationCmd  string
}

// ConfigChan is used for communicating results from concurrent parsing tasks.
type ConfigChan struct {
	StructHelper *StructGenHelper
	Destination  string
	Err          error
}

// Field represents a single field in a struct.
type Field struct {
	Name string
	Type string
	Tags string
}

func (f *Field) BuildType() string {
	return f.Type
}

// GetUsedPackageName returns the package name from the type if it's not a local type.
func (f *Field) GetUsedPackageName() string {
	if !strings.Contains(f.Type, ".") {
		return ""
	}

	name := strings.SplitN(f.Type, ".", 2)[0]
	return strings.TrimPrefix(name, "*")
}

// Import represents a Go import statement.
type Import struct {
	Name string
	Path string
}

// ToImport converts the Import struct to a valid Go import string.
func (i *Import) ToImport() string {
	if len(i.Name) == 0 {
		return i.Path
	}

	if i.Name+"\"" == filepath.Base(i.Path) {
		return i.Path
	}

	return fmt.Sprintf("%s %s", i.Name, i.Path)
}

// GetName returns the name of the import.
func (i *Import) GetName() string {
	if len(i.Name) == 0 {
		return filepath.Base(i.Path[1 : len(i.Path)-1])
	}

	return i.Name
}

// GetUsedPackageName returns the package name from the import path if it's not a local type.
func (i *Import) GetUsedPackageName() string {
	if !strings.Contains(i.Path, ".") {
		return ""
	}

	name := strings.SplitN(i.Path, ".", 2)[0]
	return strings.TrimPrefix(name, "*")
}

// StructGenHelper holds the state and metadata required to generate a builder for a struct.
type StructGenHelper struct {
	Name          string
	SrcPackage    string
	Fields        []*Field
	Imports       []*Import
	DstPackage    string
	GenerationCmd string

	// Used Internally
	MaxFieldLen  int
	MaxTypeLen   int
	UsedPackages map[string]struct{}
}

// FillDefaults populates default values for the Config.
func (c *Config) FillDefaults() (*Config, error) {
	if c == nil {
		return nil, fmt.Errorf("config is nil")
	}

	if c.Source == "" {
		return nil, fmt.Errorf("source file is required")
	}

	if c.Name == "" {
		return nil, fmt.Errorf("name is required")
	}

	// Current default is empty string. Skip Package
	if c.Destination == "" {
		c.Destination = c.Source[:len(c.Source)-len(filepath.Ext(c.Source))] + "_builder"
	}

	if c.ParserMode == "" {
		c.ParserMode = MODE_AST
	}

	return c, nil
}

// ToCommand generates the CLI command string for this config.
func (c *Config) ToCommand() (string, error) {
	if c == nil {
		return "", fmt.Errorf("config is nil")
	}

	if c.Source == "" {
		return "", fmt.Errorf("source file is required")
	}

	if c.Name == "" {
		return "", fmt.Errorf("name is required")
	}

	cmd := "buildergen"
	cmd += " --src=" + c.Source
	cmd += " --name=" + c.Name

	if c.Destination != "" {
		cmd += " --dst=" + c.Destination
	}

	if c.Package != "" {
		cmd += " --pkg=" + c.Package
	}

	if c.WithValidation {
		cmd += " --validation"
	}

	return cmd, nil
}

// Preprocess prepares the helper for source generation.
func (s *StructGenHelper) Preprocess() {
	if s.MaxFieldLen == 0 {
		for _, f := range s.Fields {
			if len(f.Name) > s.MaxFieldLen {
				s.MaxFieldLen = len(f.Name)
			}

			typeLen := f.BuildType()
			if len(typeLen) > s.MaxTypeLen {
				s.MaxTypeLen = len(typeLen)
			}
		}
	}
}

// BuildStruct generates the struct definition and methods.
func (s *StructGenHelper) BuildStruct() string {
	s.Preprocess()
	var sb strings.Builder
	if s.GenerationCmd != "" {
		sb.WriteString("// Generated using: ")
		sb.WriteString(s.GenerationCmd)
		sb.WriteString("\n")
	}
	sb.WriteString("type ")
	sb.WriteString(s.Name)
	sb.WriteString("Builder struct {\n")
	for _, field := range s.Fields {
		sb.WriteString("\t")
		sb.WriteString(field.Name)
		sb.WriteString(strings.Repeat(" ", s.MaxFieldLen-len(field.Name)+1))
		sb.WriteString(field.BuildType())
		if len(field.Tags) > 0 {
			sb.WriteString(strings.Repeat(" ", s.MaxTypeLen-len(field.Type)+1))
			sb.WriteString(field.Tags)
		}
		sb.WriteString("\n")
	}
	sb.WriteString("}\n\n")

	s.GenNewMethod(&sb)

	for _, field := range s.Fields {
		s.GenMethod(&sb, field)
	}

	s.GenBuildMethod(&sb)

	return sb.String()
}

// GenNewMethod generates the New...Builder function.
func (s *StructGenHelper) GenNewMethod(builder *strings.Builder) {
	builder.WriteString("func New")
	builder.WriteString(s.Name)
	builder.WriteString("Builder(b *")
	if s.DstPackage != s.SrcPackage {
		builder.WriteString(s.SrcPackage)
		builder.WriteString(".")
	}
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
		builder.WriteString(strings.Repeat(" ", s.MaxFieldLen-len(field.Name)+1))
		builder.WriteString("b.")
		builder.WriteString(field.Name)
		builder.WriteString(",")
	}
	builder.WriteString("\n\t}\n}\n\n")
}

// GenMethod generates a setter method for a field.
func (s *StructGenHelper) GenMethod(builder *strings.Builder, field *Field) {
	// This is a placeholder for the actual implementation logic that was in ttypes.go
}

// GenBuildMethod generates the Build method.
func (s *StructGenHelper) GenBuildMethod(builder *strings.Builder) {
	// This is a placeholder for the actual implementation logic that was in ttypes.go
}

// GetUsedPackages returns a set of all used packages.
func (s *StructGenHelper) GetUsedPackages() utils.Set[string] {
	set := make(utils.Set[string])
	for k := range s.UsedPackages {
		set.Add(k)
	}
	return set
}
