# Command Line Arguments

The command line consists of 4 different flags.

| Command Flag | Description                                     | Example                                     |
| ------------ | ----------------------------------------------- | ------------------------------------------- |
| `--src`      | Required. Source file path                      | `buildergen --src examples/test.go`         |
| `--name`     | Required. Name of the Struct to build           | `buildergen --name Person`                  |
| `--dst`      | Optional. The destination of the generated file | `buildergen --dst examples/test_builder.go` |
| `--pkg`      | Optional. The package of the generated file     | `buildergen --pkg examples`                 |
| `--validate` | Optional. Validates the source file syntax      | `buildergen --validate`                     |
| `--config`   | The configuration file used to generate         | `buildergen --config`                       |

## Command line

### Generating 1 file

```bash
buildergen --src ./test.go --name Person --dst ./test_builder.go --pkg examples
```

### Generating multiple files

```bash
buildergen --config ./config.yaml
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