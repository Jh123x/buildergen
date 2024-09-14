package cmd

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/Jh123x/buildergen/internal/consts"
	"gopkg.in/yaml.v3"
)

type BuilderGenConfig struct {
	// DefaultOutputDir is the directory to write each builder when a destination is not specified.
	DefaultOutputDir string `yaml:"default-dir"`
	// DefaultPackage is the package name to write each builder when there are no packages specified.
	DefaultPackage string    `yaml:"default-package"`
	Configs        []*Config `yaml:"configs"`
}

func ParseConfigFile(configFile string) ([]*Config, error) {
	if len(configFile) == 0 {
		return nil, consts.ErrInvalidConfigFile
	}

	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}

	builderConfig := new(BuilderGenConfig)
	if err := yaml.Unmarshal(data, &builderConfig); err != nil {
		return nil, err
	}

	configs := make([]*Config, 0, len(builderConfig.Configs))
	for _, conf := range builderConfig.Configs {
		if conf.Destination == "" && conf.Source != "" && builderConfig.DefaultOutputDir != "" {
			sourceFileName := path.Base(conf.Source)
			conf.Destination = path.Join(builderConfig.DefaultOutputDir, sourceFileName[:strings.LastIndex(sourceFileName, ".")]+consts.DEFAULT_BUILDER_SUFFIX)
		}

		if conf.Package == "" && builderConfig.DefaultPackage != "" {
			conf.Package = builderConfig.DefaultPackage
		}

		conf, err := conf.FillDefaults()
		if err != nil {
			return nil, err
		}
		fmt.Printf("%+v\n", conf)

		configs = append(configs, conf)
	}

	return configs, nil
}
