// This is the struct used for examples
package examples

import (
	"os"

	"github.com/Jh123x/buildergen/examples/nested"
	"golang.org/x/tools/imports"
)

//go:generate buildergen -src ./person.go -name Person

type Person struct {
	ID        int
	Name      string
	Email     *string // Optional field
	PhoneBook []*Contact
	MapVal    map[string]string `json:"map_val"`
	T         nested.Test
}

type Contact struct {
	Name  string
	Phone string
}

type UnRelated struct {
	importOpts *imports.Options
	otherOpts  *os.FileMode
}
