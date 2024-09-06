// Code generated by BuilderGen v0.0.7
package benchmark

import "github.com/Jh123x/buildergen/examples"

type DataBuilder struct {
	Person     *examples.Person
	SpareData  int
	SpareData2 string
	SpareData3 []*SubData
}

func NewDataBuilder(b *Data) *DataBuilder {
	if b == nil {
		return nil
	}

	return &DataBuilder{
		Person:     b.Person,
		SpareData:  b.SpareData,
		SpareData2: b.SpareData2,
		SpareData3: b.SpareData3,
	}
}

func (b *DataBuilder) WithPerson(person *examples.Person) *DataBuilder {
	b.Person = person
	return b
}

func (b *DataBuilder) WithSpareData(spareData int) *DataBuilder {
	b.SpareData = spareData
	return b
}

func (b *DataBuilder) WithSpareData2(spareData2 string) *DataBuilder {
	b.SpareData2 = spareData2
	return b
}

func (b *DataBuilder) WithSpareData3(spareData3 []*SubData) *DataBuilder {
	b.SpareData3 = spareData3
	return b
}

func (b *DataBuilder) Build() *Data {
	return &Data{
		Person:     b.Person,
		SpareData:  b.SpareData,
		SpareData2: b.SpareData2,
		SpareData3: b.SpareData3,
	}
}
