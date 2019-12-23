package main

import (
	"os"
	"path/filepath"

	"github.com/andyinabox/go-dumbphone/cmd/dumbp/cmd"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
)

const (
	version    = "0.2.0"
	configDir  = ".dumbp/"
	configName = "config"
	configExt  = "yml"
)

func main() {

	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	configPath := filepath.Join(home, configDir)

	viper.SetConfigType("yaml")
	viper.SetConfigName(configName)
	viper.AddConfigPath(configPath)
	// viper.AddConfigPath(".")

	// try reading in config
	if err := viper.ReadInConfig(); err != nil {
		// if no config found, create one from the example file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

			// try to create
			err = createConfig(configPath, configName+"."+configExt, "./config.example.yml")
			if err != nil {
				panic(err)
			}

			// try reading in config again
			err := viper.ReadInConfig()
			if err != nil {
				panic(err)
			}

			// panic for other errors
		} else {
			panic(err)
		}
	}

	app := cli.NewApp()
	app.Name = "dumbphone"
	app.Usage = "A set of tools to make your dumbphone a little smarter"
	app.Version = version
	app.Commands = []cli.Command{
		cmd.DirectionsSubcommand,
		cmd.ReaderSubcommand,
		cmd.MarkdownSubcommand,
		cmd.ConfigSubcommand,
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
