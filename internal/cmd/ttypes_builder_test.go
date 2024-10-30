// Code generated by BuilderGen v0.2.1
package cmd

import "github.com/Jh123x/buildergen/internal/consts"

type testCaseBuilder struct {
	src            string
	dst            string
	pkg            string
	name           string
	withValidation bool
	parserMode     consts.Mode
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
		withValidation: b.withValidation,
		parserMode:     b.parserMode,
		expectedConfig: b.expectedConfig,
		expectedErr:    b.expectedErr,
	}
}

func (b *testCaseBuilder) Withsrc(src string) *testCaseBuilder {
	b.src = src
	return b
}

func (b *testCaseBuilder) Withdst(dst string) *testCaseBuilder {
	b.dst = dst
	return b
}

func (b *testCaseBuilder) Withpkg(pkg string) *testCaseBuilder {
	b.pkg = pkg
	return b
}

func (b *testCaseBuilder) Withname(name string) *testCaseBuilder {
	b.name = name
	return b
}

func (b *testCaseBuilder) WithwithValidation(withValidation bool) *testCaseBuilder {
	b.withValidation = withValidation
	return b
}

func (b *testCaseBuilder) WithparserMode(parserMode consts.Mode) *testCaseBuilder {
	b.parserMode = parserMode
	return b
}

func (b *testCaseBuilder) WithexpectedConfig(expectedConfig *Config) *testCaseBuilder {
	b.expectedConfig = expectedConfig
	return b
}

func (b *testCaseBuilder) WithexpectedErr(expectedErr error) *testCaseBuilder {
	b.expectedErr = expectedErr
	return b
}

func (b *testCaseBuilder) Build() *testCase {
	return &testCase{
		src:            b.src,
		dst:            b.dst,
		pkg:            b.pkg,
		name:           b.name,
		withValidation: b.withValidation,
		parserMode:     b.parserMode,
		expectedConfig: b.expectedConfig,
		expectedErr:    b.expectedErr,
	}
}
