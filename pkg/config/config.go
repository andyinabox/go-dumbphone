package config

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/jinzhu/configor"
	"gopkg.in/yaml.v2"
)

// Config is the config structure
type Config struct {
	USB struct {
		Root          string `required:"true"`
		NotesDir      string `default:"/Notes"`
		PodcastsDir   string `default:"/Podcasts"`
		DirectionsDir string `default:"/Directions"`
		ReaderDir     string `default:"/Reading"`
	}

	Directions struct {
		GoogleMapsAPIKey string `required:"true"`
		HomeAddress      string
	}

	Notes struct {
		NotesDir string
	}
}

// Load configuration file at given path
func Load(fp string) (*Config, bool, error) {

	var exists = true
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		exists = false
	}

	c := Config{}
	err := configor.Load(&c, fp)

	return &c, exists, err
}

// Save resaves the config file
func Save(c *Config, fp string) error {

	// make directory if it doesn't exist
	p := filepath.Dir(fp)
	if _, err := os.Stat(p); os.IsNotExist(err) {
		os.MkdirAll(p, 0777)
	}

	m, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(fp)
	defer f.Close()
	if err != nil {
		return err
	}

	_, err = f.Write(m)
	if err != nil {
		return err
	}

	return nil
}

// IsRequiredErr checks if the error was triggered
// because of a missing required field
func IsRequiredErr(e error) bool {
	search := "is required, but blank"

	return strings.Contains(e.Error(), search)
}
