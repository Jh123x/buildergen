package cmd

import (
	"flag"
	"fmt"
)

// GetUsage prints the help message for BuilderGen.
func GetUsage() {
	fmt.Printf("BuilderGen is a builder code generation library to easily create builders around your struct\n")
	fmt.Printf("Usage Example: buildergen -src ./examples/test.go -name Person\n")
	flag.VisitAll(flagPrinter)
}

func flagPrinter(f *flag.Flag) {
	fmt.Printf("- %-5s: %s\n", f.Name, f.Usage)
}
