package cmd

import (
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

//go:generate buildergen -src ttypes_test.go -name testCase -dst ttypes_builder_test.go

type testCase struct {
	src  *string
	dst  *string
	pkg  *string
	name *string

	expectedConfig *Config
	expectedErr    error
}

var (
	defaultSrc    = "test_src.go"
	defaultDst    = "test_dst.go"
	defaultPkg    = "test"
	defaultName   = "TestCase"
	notGoSrc      = "not_go_ext"
	defaultConfig = &Config{
		Source:      defaultSrc,
		Destination: defaultDst,
		Package:     defaultPkg,
		Name:        defaultName,
	}
	defaultSuccessTestcase = &testCase{
		src:  &defaultSrc,
		dst:  &defaultDst,
		pkg:  &defaultPkg,
		name: &defaultName,

		expectedConfig: defaultConfig,
		expectedErr:    nil,
	}
)

func TestNewConfig(t *testing.T) {
	tests := map[string]*testCase{
		"success": NewtestCaseBuilder(defaultSuccessTestcase).Build(),
		"empty src should error": NewtestCaseBuilder(defaultSuccessTestcase).
			Withsrc(nil).
			WithexpectedErr(consts.ErrSrcNotFound).
			WithexpectedConfig(nil).
			Build(),
		"src not go file should error": NewtestCaseBuilder(defaultSuccessTestcase).
			Withsrc(&notGoSrc).
			WithexpectedErr(consts.ErrNotGoFile).
			WithexpectedConfig(nil).
			Build(),
		"empty name should error": NewtestCaseBuilder(defaultSuccessTestcase).
			Withname(nil).
			WithexpectedErr(consts.ErrNameNotFound).
			WithexpectedConfig(nil).
			Build(),
		"empty dst should return dst based on src": NewtestCaseBuilder(defaultSuccessTestcase).
			Withdst(nil).
			WithexpectedConfig(
				NewConfigBuilder(defaultConfig).
					WithDestination("test_src_builder.go").
					Build(),
			).Build(),
		"empty pkg should return default pkg": NewtestCaseBuilder(defaultSuccessTestcase).
			Withpkg(nil).
			WithexpectedConfig(
				NewConfigBuilder(defaultConfig).
					WithPackage("").
					Build(),
			).Build(),
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := NewConfig(tc.src, tc.dst, tc.pkg, tc.name)
			assert.Equal(t, tc.expectedConfig, cfg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
