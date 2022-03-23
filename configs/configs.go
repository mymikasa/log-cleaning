package configs

import (
	"bytes"
	"cleaning/pkg/file"
	_ "embed"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Settings struct {
	Logger struct {
		Version    string `yaml:"version"`
		LogPath    string `yaml:"logpath"`
		Zipdays    int    `yaml:"zipdays"`
		Normaldays int    `yaml:"normaldays"`
	}
}

var config = new(Settings)

var (
	//go:embed settings.yaml
	settings []byte
)

func init() {
	// var r io.Reader
	r := bytes.NewReader(settings)
	viper.SetConfigType("yaml")

	if err := viper.ReadConfig(r); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(config); err != nil {
		panic(err)
	}

	viper.SetConfigName("settings")
	viper.AddConfigPath("./configs")

	configFile := "./configs/settings.yaml"
	_, ok := file.IsExists(configFile)

	if !ok {
		if err := os.MkdirAll(filepath.Dir(configFile), 0766); err != nil {
			panic(err)
		}

		f, err := os.Create(configFile)
		if err != nil {
			panic(err)
		}

		defer f.Close()

		if err := viper.WriteConfig(); err != nil {
			panic(err)
		}
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(config); err != nil {
			panic(err)
		}
	})
}

func Get() Settings {
	return *config
}
