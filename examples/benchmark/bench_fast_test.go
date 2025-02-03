// The main package for running struct generation benchmarks
package benchmark

import (
	"os"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/stretchr/testify/assert"
)

var (
	fastConfig = &cmd.Config{
		Source:      "./benchmark.go",
		Destination: "./benchmark_builder.go",
		Package:     "benchmark",
		Name:        "Data",
		ParserMode:  consts.MODE_FAST,
	}

	expectedRes, _ = os.ReadFile("./benchmark_builder.go")
)

func BenchmarkCodeGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := parser.ParseBuilderFile(fastConfig)
		assert.Nil(b, err)

		assert.Equal(b, string(expectedRes), data)
	}
}

func BenchmarkCodeGenWithIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := parser.ParseBuilderFile(fastConfig)
		assert.Nil(b, err)

		file, err := os.Create(fastConfig.Destination)
		assert.Nil(b, err)

		file.WriteString(data.ToSource())
		assert.Nil(b, file.Close())
	}
}
