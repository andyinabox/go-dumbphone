package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"

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
	USB        *usb.Config        `key:"usb" desc:"Settings for managing your phone via USB"`
	Directions *directions.Config `key:"directions" desc:"Settings for the directions command"`
	Notes      *notes.Config      `key:"notes" desc:"Settings for the notes command"`
}

var configFilePath string
var config *Configuration
var configFileObject *os.File
var currentGroup interface{}

// Load should always be called initially to load
// config file or create an empty one
func Load() error {

	// get the user's home dir
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// set the full config file path
	configFilePath = fmt.Sprintf("%s/%s/%s", home, ConfigDirName, ConfigFileName)

	// read or create config
	err = getConfig(config)
	if err != nil {
		return err
	}

	return nil
}



// Save saves the config file to the filesystem
func Save() error {

	if config == nil {
		return errors.New("Configuration has not yet been loaded")
	}

	err := writeConfig(config)
	if err != nil {
		return err
	}

	return nil
}




func getConfig(c *Configuration) error {

	// check to see if the file exists
	_, err := os.Stat(configFilePath)
	if err != nil {

		// if not, create default config and write the file
		if os.IsNotExist(err) {
			c = createConfig()
			err = writeConfig(c)
			if err != nil {
				return err
			}

			// else return any other errors
		} else {
			return err
		}
		// if it exists, read in the config file
	} else {
		err = readConfig(c)
		if err != nil {
			return err
		}
	}

	return nil
}

// read config file into memory
func readConfig(c *Configuration) error {
	var c *Configuration

	// open file
	f, err := os.Open(configFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// read file into bytes
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// unmarshal bytes into struct
	err = yaml.Unmarshal(b, c)
	if err != nil {
		return err
	}

	return nil
}

func writeConfig(c *Configuration) error {
	f, err := os.Open(configFilePath)
	if err != nil {
		return err
	}
	defer f.Close()

	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	_, err = f.Write(b)
	if err != nil {
		return err
	}

	return nil
}

func createConfig() *Configuration {
	return &Configuration{
		usb.ConfigDefaults,
		directions.ConfigDefaults,
		notes.ConfigDefaults,
	}
}
