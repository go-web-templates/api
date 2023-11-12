package conf

import (
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"go.uber.org/fx"
)

var Module = fx.Options(fx.Provide(NewAppConf))

type AppConf struct {
	RunMode string
	Port    string `config:"port"`
	Data    struct {
		Database struct {
			User string `config:"user"`
			Pass string `config:"pass"`
			Name string `config:"name"`
			Host string `config:"host"`
			Port string `config:"port"`
		} `config:"database"`
		Cache struct {
			Host string `config:"host"`
			Pass string `config:"pass"`
			Port string `config:"port"`
		} `config:"cache"`
	}
}

func NewAppConf() (*AppConf, error) {
	appConf := &AppConf{}

	config.WithOptions(config.ParseEnv, func(o *config.Options) {
		o.DecoderConfig.TagName = "config"
	})

	config.AddDriver(yaml.Driver)

	confFilename := "./app-conf-dev.yml"
	if RunMode == RUN_MODE_RELEASE {
		confFilename = "./app-conf.yml"
	}

	if err := config.LoadFiles(confFilename); err != nil {
		return appConf, err
	}

	if err := config.Decode(&appConf); err != nil {
		return appConf, err
	}

	appConf.RunMode = RunMode
	return appConf, nil
}
