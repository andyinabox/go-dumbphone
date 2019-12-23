package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanity(t *testing.T) {
	fmt.Println("You appear to be sane")
}

func TestCreateConfig(t *testing.T) {
	assert := assert.New(t)
	testPath := "../../test/config/"
	testFileName := "config.yml"
	filePath := filepath.Join(testPath, testFileName)
	configPath := "../../config.example.yml"

	{
		err := createConfig(testPath, testFileName, configPath)
		assert.Nil(err, "Expected no error executing createConfig")

		_, err = os.Stat(filePath)
		assert.False(os.IsNotExist(err), "File does not exist")
		assert.Nil(err, "Some other file error")
	}

	os.RemoveAll(testPath)
}
