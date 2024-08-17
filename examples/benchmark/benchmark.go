// This is the code used to benchmark the performance of BuilderGen
package benchmark

import "github.com/Jh123x/buildergen/examples"

//go:generate buildergen -src ./benchmark.go -name Data

type Data struct {
	Person     *examples.Person
	SpareData  int
	SpareData2 string
	SpareData3 []*SubData
}

type SubData struct {
	SubData SubSubData
}

type SubSubData struct {
	SomeData  string
	SomeData2 *SubData
}
