package cmd

import (
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

//go:generate buildergen -src ttypes_test.go -name testCase -dst ttypes_builder_test.go

type testCase struct {
	src            string
	dst            string
	pkg            string
	name           string
	withValidation bool

	expectedConfig *Config
	expectedErr    error
}

const (
	defaultSrc        = "test_src.go"
	defaultDst        = "test_dst.go"
	defaultPkg        = "test"
	defaultName       = "TestCase"
	notGoSrc          = "not_go_ext"
	defaultValidation = false
)

var (
	defaultConfig = &Config{
		Source:         defaultSrc,
		Destination:    defaultDst,
		Package:        defaultPkg,
		Name:           defaultName,
		WithValidation: defaultValidation,
	}
	defaultSuccessTestCase = &testCase{
		src:            defaultSrc,
		dst:            defaultDst,
		pkg:            defaultPkg,
		name:           defaultName,
		withValidation: defaultValidation,

		expectedConfig: defaultConfig,
		expectedErr:    nil,
	}
)

func TestNewConfig(t *testing.T) {
	tests := map[string]*testCase{
		"success": NewtestCaseBuilder(defaultSuccessTestCase).Build(),
		"empty src should error": NewtestCaseBuilder(defaultSuccessTestCase).
			Withsrc("").
			WithexpectedErr(consts.ErrSrcNotFound).
			WithexpectedConfig(nil).
			Build(),
		"src not go file should error": NewtestCaseBuilder(defaultSuccessTestCase).
			Withsrc(notGoSrc).
			WithexpectedErr(consts.ErrNotGoFile).
			WithexpectedConfig(nil).
			Build(),
		"empty name should error": NewtestCaseBuilder(defaultSuccessTestCase).
			Withname("").
			WithexpectedErr(consts.ErrNameNotFound).
			WithexpectedConfig(nil).
			Build(),
		"empty dst should return dst based on src": NewtestCaseBuilder(defaultSuccessTestCase).
			Withdst("").
			WithexpectedConfig(
				NewConfigBuilder(defaultConfig).
					WithDestination("test_src_builder.go").
					Build(),
			).Build(),
		"empty pkg should return default pkg": NewtestCaseBuilder(defaultSuccessTestCase).
			Withpkg("").
			WithexpectedConfig(
				NewConfigBuilder(defaultConfig).
					WithPackage(consts.EMPTY_STR).
					Build(),
			).Build(),
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			cfg, err := NewConfig(tc.src, tc.dst, tc.pkg, tc.name, tc.withValidation)
			assert.Equal(t, tc.expectedConfig, cfg)
			assert.Equal(t, tc.expectedErr, err)
		})
	}
}
