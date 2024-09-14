package cmd

import (
	"os"

	"github.com/Jh123x/buildergen/internal/consts"
	"gopkg.in/yaml.v3"
)

type BuilderGenConfig struct {
	Configs []*Config `yaml:"configs"`
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
		conf, err := conf.FillDefaults()
		if err != nil {
			return nil, err
		}

		configs = append(configs, conf)
	}

	return configs, nil
}
