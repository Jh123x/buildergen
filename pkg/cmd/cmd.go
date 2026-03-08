package cmd

import (
	"flag"
	"os"

	"github.com/Jh123x/buildergen/pkg/consts"
	"gopkg.in/yaml.v3"
)

const (
	UsageFormat = "BuilderGen is a builder code generation library to easily create builders around your struct\nUsage Example:\nSingle files: `buildergen -src ./examples/test.go -name Person`\nConfig files: `buildergen -config config.yml`"
)

type BuilderGenConfig struct {
	Configs []*Config `yaml:"configs"`
}

var _ BuilderGenConfig = BuilderGenConfig{}

var flagPrinterGen func(PrinterFn) func(f *flag.Flag)

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

func GetUsage(formatPrinter PrinterFn) {
	formatPrinter(UsageFormat)
	flag.VisitAll(flagPrinterGen(formatPrinter))
}

func init() {
	flagPrinterGen = func(formatPrinter PrinterFn) func(f *flag.Flag) {
		return func(f *flag.Flag) {
			formatPrinter("- %-5s: %s\n", f.Name, f.Usage)
		}
	}
}
