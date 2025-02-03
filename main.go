// BuilderGen is a code generation tool to generate builder structs based on your structs.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/Jh123x/buildergen/internal/utils"
)

func logWrapper(s string, a ...any) (int, error) {
	log.Printf(s, a...)
	return 0, nil
}

func main() {
	var (
		src            = flag.String("src", consts.EMPTY_STR, "the source file path")
		name           = flag.String("name", consts.EMPTY_STR, "the name of the struct")
		dest           = flag.String("dst", consts.EMPTY_STR, "the destination file path, default: {src_dir}/{src}_builder.go")
		pkg            = flag.String("pkg", consts.EMPTY_STR, "the package name of the generated file, default: {src pkg}")
		withValidation = flag.Bool("validate", false, "validate the syntax of the original file, default: false")
		astMode        = flag.String("mode", string(consts.MODE_AST), "the parser mode")

		configFile = flag.String("config", consts.EMPTY_STR, "the config file for buildergen")
	)

	flag.Parse()

	if !utils.IsNilOrEmpty(configFile) {
		parser.ParseAndWriteBuilderFile(*configFile)
		return
	}

	if !utils.IsNilOrEmpty(src) {
		parser.ParseCommand(*src, *dest, *pkg, *name, *withValidation, *astMode, logWrapper)
		return
	}

	cmd.GetUsage(fmt.Printf)
	os.Exit(1)
}
