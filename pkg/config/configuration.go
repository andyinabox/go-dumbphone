package config

import (
	"errors"
	"io/ioutil"
	"os"
	"path"

	"github.com/jinzhu/copier"
	"gopkg.in/yaml.v2"

	"github.com/andyinabox/go-dumbphone/internal/utils"
	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/andyinabox/go-dumbphone/pkg/notes"
	"github.com/andyinabox/go-dumbphone/pkg/usb"
)

// Configuration is the App-level configuration struct
type Configuration struct {
	USB        *usb.Config        `key:"usb" desc:"Settings for managing your phone via USB"`
	Directions *directions.Config `key:"directions" desc:"Settings for the directions command"`
	Notes      *notes.Config      `key:"notes" desc:"Settings for the notes command"`
}

// ConfigFilePath path to the config file
var ConfigFilePath string

var config *Configuration

var groupMap map[string]interface{}

// Load should always be called initially to load
// config file or create an empty one
func Load(filePath string) error {

	ConfigFilePath = filePath

	// read or create config
	config, err := getConfig()
	if err != nil {
		return err
	}

	groupMap, err = utils.GetTagMapReverse(config, "key")
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

	err := writeConfig(config, ConfigFilePath)
	if err != nil {
		return err
	}

	return nil
}

// GetGroup get a config group
func Get(groupName string) (interface{}, error) {

	if config == nil {
		return nil, errors.New("Must initialize with Load first")
	}

	err := validateGroupName(groupName)
	if err != nil {
		return nil, err
	}
	return groupMap[groupName], nil
}

func validateGroupName(s string) error {
	found := false

	for k := range groupMap {
		if s == k {
			found = true
		}
	}

	if !found {
		return errors.New("Not a valid config group name")
	}

	return nil
}

func getConfig() (*Configuration, error) {
	var c *Configuration

	// check to see if the file exists
	_, err := os.Stat(ConfigFilePath)
	if err != nil {

		// if not, create default config and write the file
		if os.IsNotExist(err) {
			c = createConfig()
			err = writeConfig(c, ConfigFilePath)
			if err != nil {
				return nil, err
			}

			// else return any other errors
		} else {
			return nil, err
		}
		// if it exists, read in the config file
	} else {
		err = readConfig(c, ConfigFilePath)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

// read config file into memory
func readConfig(c *Configuration, filePath string) error {

	// open file
	f, err := os.Open(filePath)
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

func writeConfig(c *Configuration, filePath string) error {
	var f *os.File

	// make sure path exists
	p := path.Dir(filePath)
	err := os.MkdirAll(p, 0777)
	if err != nil {
		return err
	}

	f, err = os.Create(filePath)
	defer f.Close()

	if err != nil {
		return err
	}

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

	u := usb.Config{}
	d := directions.Config{}
	n := notes.Config{}

	copier.Copy(&u, usb.ConfigDefaults)
	copier.Copy(&d, directions.ConfigDefaults)
	copier.Copy(&n, notes.ConfigDefaults)

	return &Configuration{&u, &d, &n}
}
