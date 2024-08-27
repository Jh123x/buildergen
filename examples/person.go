// This is the struct used for examples
package examples

import "github.com/Jh123x/buildergen/examples/nested"

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
