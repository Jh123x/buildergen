// BuilderGen is a code generation tool to generate builder structs based on your structs.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"sync"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/parser"
	"github.com/Jh123x/buildergen/internal/utils"
)

var (
	src            = flag.String("src", consts.EMPTY_STR, "the source file path")
	name           = flag.String("name", consts.EMPTY_STR, "the name of the struct")
	dest           = flag.String("dst", consts.EMPTY_STR, "the destination file path, default: {src_dir}/{src}_builder.go")
	pkg            = flag.String("pkg", consts.EMPTY_STR, "the package name of the generated file, default: {src pkg}")
	withValidation = flag.Bool("validate", false, "validate the syntax of the original file, default: false")

	configFile = flag.String("config", consts.EMPTY_STR, "the config file for buildergen")
)

func main() {
	flag.Parse()

	if !utils.IsNilOrEmpty(configFile) {
		configs, err := cmd.ParseConfigFile(*configFile)
		if err != nil {
			cmd.GetUsage(fmt.Printf)
			fmt.Printf("Error parsing config file: %s", err.Error())
			return
		}

		var wg sync.WaitGroup
		wg.Add(len(configs))
		defer wg.Wait()

		for _, conf := range configs {
			go func() {
				defer wg.Done()
				if path.Dir(conf.Source) != path.Dir(conf.Destination) {
					fmt.Printf("[%s::%s] dest in different path from destination is currently not supported\n", conf.Source, conf.Name)
					return
				}

				if err := generateFile(conf); err != nil {
					panic(err)
				}
			}()
		}

		return
	}

	if !utils.IsNilOrEmpty(src) {
		config, err := cmd.NewConfig(*src, *dest, *pkg, *name, *withValidation)
		if err != nil {
			cmd.GetUsage(fmt.Printf)
			fmt.Printf("Error parsing config file: %s", err.Error())
			return
		}

		generateFile(config)
		return
	}

	cmd.GetUsage(fmt.Printf)
	os.Exit(1)
}

func generateFile(config *cmd.Config) error {
	builderSrc, err := parser.ParseBuilderFile(config)
	if err != nil {
		return err
	}

	file, err := os.Create(config.Destination)
	if err != nil {
		return err
	}

	file.WriteString(builderSrc)
	if err := file.Close(); err != nil {
		log.Println(err.Error())
	}

	return nil
}
