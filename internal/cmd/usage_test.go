package cmd

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/stretchr/testify/assert"
)

func TestGetUsage(t *testing.T) {
	tests := map[string]struct {
		flagDefineFn   func()
		expectedOutput string
	}{
		"no flags": {
			flagDefineFn:   func() {},
			expectedOutput: usageFormat,
		},
		"with flags": {
			flagDefineFn: func() {
				_ = flag.String("test", consts.EMPTY_STR, "test usage")
			},
			expectedOutput: usageFormat + "- test : test usage\n",
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			defer func() { assert.Nil(t, recover()) }()

			// Reset flag vars
			flag.CommandLine = flag.NewFlagSet(os.Args[0], flag.ExitOnError)
			tc.flagDefineFn()

			output := testPrinterFn(GetUsage)
			assert.Equal(t, tc.expectedOutput, output)
		})
	}
}

func testPrinterFn(fn func(PrinterFn)) string {
	builder := strings.Builder{}

	data := func(v string, elems ...any) (int, error) {
		builder.WriteString(fmt.Sprintf(v, elems...))
		return 0, nil
	}

	fn(data)
	return builder.String()
}
