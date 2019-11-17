package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/andyinabox/go-dumbphone/pkg/notes"
	"github.com/andyinabox/go-dumbphone/pkg/usb"
)

const (
	// ConfigFileName the name of the config file
	ConfigFileName string = "config.yaml"
	// ConfigDirName the name of the dir in home folder
	ConfigDirName string = ".dumbp"
)

// Configuration is the App-level configuration struct
type Configuration struct {
	USB        *usb.Config        `desc: "Settings for managing your phone via USB"`
	Directions *directions.Config `desc: "Settings for the directions command"`
	Notes      *notes.Config      `desc: "Settings for the notes command"`
}

var configDir string
var configFile string
var config *Configuration
var configFileObject *os.File

func Load() error {

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configDir = fmt.Sprintf("%s/%s", home, ConfigDirName)
	configFile = fmt.Sprintf("%s/%s", configDir, ConfigFileName)

	_, err = os.Stat(configFile)

	if os.IsNotExist(err) {
		config = getDefaultConfig()
		WriteConfig()
	} else if err {
		return err
	}

	configFileObject = os.Open(configFile)

	return nil
}

func Close() {
	if configFileObject != nil {
		configFileObject.Close()
	}
}

func getDefaultConfig() *Configuration {
	return &Configuration{
		usb.ConfigDefaults,
		directions.ConfigDefaults,
		notes.ConfigDefaults,
	}
}

func GetGroup(key string) interface{} {
	return nil
}

func GetValue(key1 string, key2 string) interface{} {
	return nil
}

func SetValue(key1 string, key2 string, value interface{}) bool {
	return nil
}

func WriteConfig() error {

	if config == nil {
		return errors.New("Configuration has not yet been loaded")
	}

	return nil
}
