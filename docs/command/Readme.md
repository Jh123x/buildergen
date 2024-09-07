# Command Line Arguments

The command line consists of 4 different flags.

| Command Flag | Description                                     | Example                                   |
| ------------ | ----------------------------------------------- | ----------------------------------------- |
| `--src`      | Required. Source file path                      | `buildgen --src examples/test.go`         |
| `--name`     | Required. Name of the Struct to build           | `buildgen --name Person`                  |
| `--dst`      | Optional. The destination of the generated file | `buildgen --dst examples/test_builder.go` |
| `--pkg`      | Optional. The package of the generated file     | `buildgen --pkg examples`                 |
| `--validate` | Optional. Validates the source file syntax      | `buildergen --validate`                   |

## Command line

```bash
buildgen --src ./test.go --name Person --dst ./test_builder.go --pkg examples
```

## Go Generate

```go

//go:generate buildergen -src ./test.go -name Person

type Person struct {
	ID        int
	Name      string
	Email     *string // Optional field
	PhoneBook []*Contact
	MapVal    map[string]string `json:"map_val"`
	T         nested.Test
}

```