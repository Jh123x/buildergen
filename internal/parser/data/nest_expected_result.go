// Code generated by BuilderGen
package nested

type TestBuilder struct {
	Val string
}

func NewTestBuilder(b *Test) *TestBuilder {
	if b == nil {
		return nil
	}

	return &TestBuilder{
		Val: b.Val,
	}
}

func (b *TestBuilder) WithVal(val string) *TestBuilder {
	b.Val = val
	return b
}

func (b *TestBuilder) Build() *Test {
	return &Test{
		Val: b.Val,
	}
}