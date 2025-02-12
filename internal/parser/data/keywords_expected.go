// Code generated by BuilderGen v0.3.0
package data

type StructBuilder struct {
	Package string
	Go      string
	Func    string
	PackagE string
}

func NewStructBuilder(b *Struct) *StructBuilder {
	if b == nil {
		return nil
	}

	return &StructBuilder{
		Package: b.Package,
		Go:      b.Go,
		Func:    b.Func,
		PackagE: b.PackagE,
	}
}

func (b *StructBuilder) WithPackage(package_ string) *StructBuilder {
	b.Package = package_
	return b
}

func (b *StructBuilder) WithGo(go_ string) *StructBuilder {
	b.Go = go_
	return b
}

func (b *StructBuilder) WithFunc(func_ string) *StructBuilder {
	b.Func = func_
	return b
}

func (b *StructBuilder) WithPackagE(packagE string) *StructBuilder {
	b.PackagE = packagE
	return b
}

func (b *StructBuilder) Build() *Struct {
	return &Struct{
		Package: b.Package,
		Go:      b.Go,
		Func:    b.Func,
		PackagE: b.PackagE,
	}
}
