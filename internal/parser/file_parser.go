package parser

import (
	"fmt"
	"path"

	"github.com/Jh123x/buildergen/internal/cmd"
	"github.com/Jh123x/buildergen/internal/consts"
	"github.com/Jh123x/buildergen/internal/generation"
	"github.com/Jh123x/buildergen/internal/utils"
	"github.com/Jh123x/buildergen/internal/writer"
)

func ParseAndWriteBuilderFile(configPath string) {
	configs, err := cmd.ParseConfigFile(configPath)
	if err != nil {
		cmd.GetUsage(fmt.Printf)
		fmt.Printf("Error parsing config file: %s", err.Error())
		return
	}

	cfgChannel := make(chan cmd.ConfigChan, len(configs))
	for _, conf := range configs {
		go func() {
			if path.Dir(conf.Source) != path.Dir(conf.Destination) {
				cfgChannel <- cmd.ConfigChan{Err: fmt.Errorf("[%s::%s] dest in different path from destination is currently not supported", conf.Source, conf.Name)}
				return
			}

			res, err := ParseBuilderFile(conf)
			cfgChannel <- cmd.ConfigChan{
				StructHelper: res,
				Destination:  conf.Destination,
				Err:          err,
			}
		}()
	}

	// Collect the result and write to file
	mapperData := make(map[string][]cmd.ConfigChan, len(configs))

	for i := 0; i < len(configs); i++ {
		res := <-cfgChannel
		if res.Err != nil {
			fmt.Printf("%s\n", res.Err.Error())
			continue
		}
		mapperData[res.Destination] = append(mapperData[res.Destination], res)
	}

	for filePath, cfgs := range mapperData {
		if err := writer.MultiFileWrite(
			filePath,
			utils.Map(
				cfgs,
				func(cfg cmd.ConfigChan) *generation.StructGenHelper {
					return cfg.StructHelper
				},
			)...,
		); err != nil {
			fmt.Printf("%s\n", err.Error())
			return
		}
	}
}

func ParseCommand(src, dest, pkg, name string, withValidation bool, astMode string, logWrapper cmd.PrinterFn) {
	config, err := cmd.NewConfig(src, dest, pkg, name, withValidation, consts.Mode(astMode))
	if err != nil {
		cmd.GetUsage(logWrapper)
		logWrapper("Error parsing config file: %s", err.Error())
		return
	}

	builderSrc, err := ParseBuilderFile(config)
	if err != nil {
		cmd.GetUsage(logWrapper)
		logWrapper("Error parsing config file: %s", err.Error())
		return
	}

	if err := writer.WriteToSingleFile(config.Destination, builderSrc.ToSource()); err != nil {
		logWrapper("Error parsing config file: %s", err.Error())
		return
	}
}
