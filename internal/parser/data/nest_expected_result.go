// Code generated by BuilderGen v0.3.0
package data

import "os"

type TestBuilder struct {
	Val          string
	ImportedType *os.FileMode
}

func NewTestBuilder(b *Test) *TestBuilder {
	if b == nil {
		return nil
	}

	return &TestBuilder{
		Val:          b.Val,
		ImportedType: b.ImportedType,
	}
}

func (b *TestBuilder) WithVal(val string) *TestBuilder {
	b.Val = val
	return b
}

func (b *TestBuilder) WithImportedType(importedType *os.FileMode) *TestBuilder {
	b.ImportedType = importedType
	return b
}

func (b *TestBuilder) Build() *Test {
	return &Test{
		Val:          b.Val,
		ImportedType: b.ImportedType,
	}
}
