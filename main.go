package main

import (
	"cleaning/internal/clean"
	"cleaning/pkg/file"
	"cleaning/pkg/logging"
	"flag"
	"fmt"
	"os"
	"path"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

type exitCode struct{ Code int }

type Settings struct {
	Logger struct {
		Version    string `yaml:"version"`
		LogPath    string `yaml:"logpath"`
		Zipdays    int    `yaml:"zipdays"`
		Normaldays int    `yaml:"normaldays"`
	}
}

var config = new(Settings)

func Cron() {
	c := cron.New()
	_, err := c.AddFunc("20 11 * * *", run)

	if err != nil {
		return
	}
	c.Start()
	select {}
}

func run() {
	filepath := config.Logger.LogPath
	filenames, err := file.GetAllFileName(filepath)

	if err != nil {
		logging.Panic(err)
	}

	for _, fpath := range filenames {
		name := path.Base(fpath)
		h, err := clean.NewFileHandler(name, fpath)
		if err != nil {
			logging.Panic(err)
		}
		h.Clean(config.Logger.Zipdays, config.Logger.Normaldays)
	}
}

func main() {

	configPath := flag.String("config-dir", ".", "Directory that contains the configuration file")
	flag.Parse()

	viper.SetConfigName("settings")
	viper.AddConfigPath(*configPath)
	viper.SetConfigType("yaml")
	fmt.Fprintln(os.Stderr, "Reading configuration from", *configPath)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Fprintln(os.Stderr, "Failed reading configuration:", err.Error())
		panic(exitCode{1})
	}
	// logging.MustSetup(constant.IsDebug)
	m := logging.WithIsDebug(false)

	logging.MustSetup(m)
	viper.WatchConfig()
	if err := viper.Unmarshal(config); err != nil {
		panic(exitCode{1})
	}
	Cron()
}
