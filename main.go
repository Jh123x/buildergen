// BuilderGen is a code generation tool to generate builder structs based on your structs.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/Jh123x/buildergen/internal/utils"
)

var (
	src  = flag.String("src", "", "[required] the source file path")
	name = flag.String("name", "", "[required] the name of the struct")
	dest = flag.String("dst", "", "[optional] the destination file path, default: {src_dir}/{src}_builder.go")
	pkg  = flag.String("pkg", "", "[optional] the package name of the generated file, default: {src pkg}")
)

func main() {
	flag.Parse()

	if utils.IsNilOrEmpty(src) {
		cmd.GetUsage(fmt.Printf)
		os.Exit(1)
		return
	}

	config, err := cmd.NewConfig(src, dest, pkg, name)
	if err != nil {
		cmd.GetUsage(fmt.Printf)
		return
	}

	builderSrc, err := parser.ParseBuilderFile(config)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(config.Destination)
	if err != nil {
		panic(err)
	}

	file.WriteString(builderSrc)
	if err := file.Close(); err != nil {
		log.Println(err.Error())
	}
}
