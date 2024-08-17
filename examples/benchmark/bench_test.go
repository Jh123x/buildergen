// The main package for running struct generation benchmarks
package benchmark

import (
	"os"
	"testing"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/stretchr/testify/assert"
)

var config = &cmd.Config{
	Source:      "./benchmark.go",
	Destination: "./benchmark_builder.go",
	Package:     "benchmark",
	Name:        "Data",
}

func BenchmarkCodeGen(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := parser.ParseBuilderFile(config)
		assert.Nil(b, err)

		file, err := os.Create(config.Destination)
		assert.Nil(b, err)

		file.WriteString(data)
		assert.Nil(b, file.Close())
	}
}
