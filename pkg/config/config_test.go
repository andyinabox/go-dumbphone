package config

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const testConfig = "../../test/config/config.yaml"

func tearDown() {
	os.Remove(testConfig)
}

func TestIsRequiredErr(t *testing.T) {
	assert := assert.New(t)
	assert.True(IsRequiredErr(errors.New("Field is required, but blank")))
}

func TestConfig(t *testing.T) {
	assert := assert.New(t)

	{
		c, exists, err := Load(testConfig)
		assert.False(exists, "The config file should not exist yet")
		assert.NotNil(err, "There should be an error because of missing required fields")

		err = Save(c, testConfig)
		assert.Nil(err, "There should not be an error saving the config")
	}

	{
		c, exists, err := Load(testConfig)
		assert.True(exists, "The config file should now exist")
		assert.NotNil(err, "There should be an error because of missing required fields")

		c.USB.Root = "/Volumes/DUMBPHONE"
		c.Directions.GoogleMapsAPIKey = "123456"

		err = Save(c, testConfig)
		assert.Nil(err, "There should not be an error saving the config")
	}

	{
		_, exists, err := Load(testConfig)
		assert.True(exists, "The config file should now exist")
		assert.Nil(err, "Now there should be no error because the required fields are filled")
	}

	tearDown()
}
