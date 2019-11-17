package config

import (
	"fmt"
	"os"
	"yaml"

	"github.com/spf13/viper"
)

const (
	ConfigFileName string = "config"
	ConfigDirName  string = ".dumbp"
)

type Configuration struct {
	USBVolume  `desc: "Settings for managing your phone via USB"`
	Directions `desc: "Settings for the directions command"`
	Notes      `desc: "Settings for the notes command"`
}

type Volume struct {
	Root          string `desc:"Path to your phone when plugged into USB"`
	PodcastsDir   string `desc:"Podcasts folder on your phone"`
	NotesDir      string `desc:"Notes folder on your phone"`
	ReaderDir     string `desc:"Reading folder on your phone"`
	DirectionsDir string `desc:"Directions folder on your phone"`
}

type Directions struct {
	GoogleAPIKey string `desc:"Google Maps API Key"`
	HomeAddress  string `desc:"Default starting address for directions`
}

type Notes struct {
	NotesDir string `desc:"Notes folder on your computer"`
}

var ConfigDir string = fmt.Sprintf("%s/%s", os.UserHomeDir(), ConfigDirName)
var defaults = &Configuration{
	&Volume{
		"",
		"/Podcasts",
		"/Notes",
		"/Reading",
		"/Directions",
	},
	&Directions{
		"",
		"",
	},
	&Notes{
		"",
	},
}

func GetConfig() error {
	v := viper.GetViper()

	v.SetConfigName(ConfigFileName) // name of config file (without extension)
	v.AddConfigPath(ConfigDir)      // path to look for the config file in

	err := v.ReadInConfig() // Find and read the config file

	if err := viper.ReadInConfig(); err != nil {
		// config file isn't found, let's create a default config
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			err = Create()

			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	// unmarshall

	return nil
}

func Create() error {
	// save config as file
	v := viper.GetViper()

	c, err := yaml.Marshall(defaults)
	file := os.Create(fmt.Sprintf("%s/%s"), ConfigDir, configFileName)

	return nil
}
