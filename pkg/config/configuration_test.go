package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/andyinabox/go-dumbphone/pkg/directions"
	"github.com/andyinabox/go-dumbphone/pkg/notes"
	"github.com/andyinabox/go-dumbphone/pkg/usb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const configPath = "../../test/config/test.yaml"

func tearDown() {
	_ = os.RemoveAll(filepath.Dir(configPath))
}

func TestCreateConfig(t *testing.T) {
	assert := assert.New(t)

	var expected = &Configuration{
		usb.ConfigDefaults,
		directions.ConfigDefaults,
		notes.ConfigDefaults,
	}

	create := createConfig()

	assert.Equal(expected, create)

}

func TestReadConfig(t *testing.T) {
	assert := assert.New(t)
	// require := require.New(t)

	config := createConfig()

	var expected = &Configuration{
		&usb.Config{
			"/Volumes/DUMBPHONE",
			"/Podcasts",
			"/Notes",
			"/Reading",
			"/Directions",
		},
		&directions.Config{
			"1234567",
			"12345 Somewhere St. Springfield, IL",
		},
		&notes.Config{
			"~/Dropbox/Writing/Notes",
		},
	}

	err := readConfig(config, "../../test/data/config.yaml")
	assert.Nil(err)
	assert.Equal(config, expected)

	tearDown()
}

func TestWriteConfig(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	config := createConfig()

	err := writeConfig(config, configPath)
	assert.Nil(err)

	_, err = os.Stat(configPath)
	assert.Nil(err, "The written file should exist")

	expected, err := os.Open("../../test/data/empty_config.yaml")
	expectedBytes, err := ioutil.ReadAll(expected)
	expected.Close()
	require.Nil(err)

	actual, err := os.Open(configPath)
	actualBytes, err := ioutil.ReadAll(actual)
	actual.Close()
	require.Nil(err)

	assert.Equal(string(expectedBytes), string(actualBytes), "The written file should be the same as the expected file")

	err = readConfig(config, "../../test/data/config.yaml")
	require.Nil(err)

	err = writeConfig(config, configPath)
	assert.Nil(err, "Writing config should still work after file has been created")

	expected, err = os.Open("../../test/data/config.yaml")
	expectedBytes, err = ioutil.ReadAll(expected)
	expected.Close()
	require.Nil(err)

	actual, err = os.Open(configPath)
	actualBytes, err = ioutil.ReadAll(actual)
	actual.Close()
	require.Nil(err)

	assert.Equal(expectedBytes, actualBytes, "The updated written file should match config.yaml fixture")

	tearDown()
}

func TestLoad(t *testing.T) {
	assert := assert.New(t)

	_, err := Get("usb")
	assert.NotNil(err, "Attempting to get a group before loading should throw an error")

	err = Load(configPath)
	assert.Nil(err)

	usbConfig, err := Get("usb")
	assert.Nil(err)

	fmt.Sprintf("%v", usbConfig)

	err = Save()
	assert.Nil(err)

	tearDown()
}
