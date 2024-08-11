# BuilderGen

BuilderGen is a tool used for generating builders for Golang `structs`.

- [Commands](./docs/command/Readme.md "Documentation for Commands")
- [Usage](./docs/usage/Readme.md "Documentation for Usage")

## QuickStart

### Step 1: Install this package

```bash
go install github.com/Jh123x/buildergen@latest
```

Install this package start using it

### Step 2: Use the package

This is just 1 out of 2 ways to use this package.
For the full list of usage visit docs [here](./docs/usage/Readme.md "Documentation for usage").
Write the go generate comment as shown in the example below.

```go
package examples

import "github.com/Jh123x/buildergen/examples/nested"

//go:generate buildergen -src=./test.go -name Person

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
```

### Step 3: Using the builder

After running the go generate, you can use the builder similar to what is shown below.

```go

var defaultPerson = &Person{
	ID: 1,
	Name: "John",
	Email: nil,
}

...
func TestXXX(t *testing.T){
	clonedPerson := NewPersonBuilder(defaultPerson).WithID(12).WithName("Johnny").Build() // ID and Name changes
	...
	// Use clonedPerson
}
```
