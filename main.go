package main

import (
	"flag"
	"os"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/Jh123x/buildergen/internal/utils"
)

func main() {
	src := flag.String("src", "", "[required] the source file path")
	name := flag.String("name", "", "[required] the name of the struct")
	dest := flag.String("dst", "", "[optional] the destination file path, default: {src_dir}/{src}_builder.go")
	pkg := flag.String("pkg", "", "[optional] the package name of the generated file, default: {src pkg}")
	flag.Parse()

	if utils.IsNilOrEmpty(src) {
		cmd.GetUsage()
		os.Exit(1)
		return
	}

	config, err := cmd.NewConfig(src, dest, pkg, name)
	if err != nil {
		cmd.GetUsage()
		return
	}

	if err := parser.ParseBuilderFile(config); err != nil {
		panic(err)
	}
}
