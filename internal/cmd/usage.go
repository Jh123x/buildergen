package cmd

import (
	"flag"
)

const usageFormat = "BuilderGen is a builder code generation library to easily create builders around your struct\n" +
	"Usage Example:\nSingle files: `buildergen -src ./examples/test.go -name Person`\n" +
	"Config files: `buildergen -config config.yml`"

// GetUsage prints the help message for BuilderGen.
func GetUsage(formatPrinter PrinterFn) {
	formatPrinter(usageFormat)
	flag.VisitAll(flagPrinterGen(formatPrinter))
}

func flagPrinterGen(formatPrinter PrinterFn) func(f *flag.Flag) {
	return func(f *flag.Flag) {
		formatPrinter("- %-5s: %s\n", f.Name, f.Usage)
	}
}
