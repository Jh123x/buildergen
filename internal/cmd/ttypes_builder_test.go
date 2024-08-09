// Code generated by BuilderGen
package cmd

type testCaseBuilder struct {
	src            *string
	dst            *string
	pkg            *string
	name           *string
	expectedConfig *Config
	expectedErr    error
}

func NewtestCaseBuilder(b *testCase) *testCaseBuilder {
	if b == nil {
		return nil
	}

	return &testCaseBuilder{
		src:            b.src,
		dst:            b.dst,
		pkg:            b.pkg,
		name:           b.name,
		expectedConfig: b.expectedConfig,
		expectedErr:    b.expectedErr,
	}
}

func (b *testCaseBuilder) Withsrc(src *string) *testCaseBuilder {
	b.src = src
	return b
}

func (b *testCaseBuilder) Withdst(dst *string) *testCaseBuilder {
	b.dst = dst
	return b
}

func (b *testCaseBuilder) Withpkg(pkg *string) *testCaseBuilder {
	b.pkg = pkg
	return b
}

func (b *testCaseBuilder) Withname(name *string) *testCaseBuilder {
	b.name = name
	return b
}

func (b *testCaseBuilder) WithexpectedConfig(expectedconfig *Config) *testCaseBuilder {
	b.expectedConfig = expectedconfig
	return b
}

func (b *testCaseBuilder) WithexpectedErr(expectederr error) *testCaseBuilder {
	b.expectedErr = expectederr
	return b
}

func (b *testCaseBuilder) Build() *testCase {
	return &testCase{
		src:            b.src,
		dst:            b.dst,
		pkg:            b.pkg,
		name:           b.name,
		expectedConfig: b.expectedConfig,
		expectedErr:    b.expectedErr,
	}
}