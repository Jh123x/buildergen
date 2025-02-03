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

var astConfig = &cmd.Config{
	Source:      "./benchmark.go",
	Destination: "./benchmark_builder.go",
	Package:     "benchmark",
	Name:        "Data",
	ParserMode:  consts.MODE_AST,
}

func BenchmarkASTCodeGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := parser.ParseBuilderFile(astConfig)
		assert.Nil(b, err)

		assert.Equal(b, string(expectedRes), data)
	}
}

func BenchmarkASTCodeGenWithIO(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := parser.ParseBuilderFile(astConfig)
		assert.Nil(b, err)

		file, err := os.Create(astConfig.Destination)
		assert.Nil(b, err)

		file.WriteString(data.ToSource())
		assert.Nil(b, file.Close())
	}
}
