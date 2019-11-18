package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/stretchr/testify/require"
)

type Test struct {
	IntField    int    `key:"intField" label:"Integer field"`
	StringField string `key:"stringField" label:"String field"`
}

var test = &Test{
	22,
	"Value",
}

// normal tag map using "label" key
var expectedMap = map[interface{}]string{
	22:      "Integer field",
	"Value": "String field",
}

var expectedMapReverse = map[string]interface{}{
	"intField":    22,
	"stringField": "Value",
}

func TestGetTag(t *testing.T) {
	assert := assert.New(t)
	// require := require.New(t)

	expectedIntFieldLabel := "Integer field"
	intFieldLabel, err := GetTag(test, "IntField", "label")

	assert.Nil(err)
	assert.Equal(expectedIntFieldLabel, intFieldLabel)

	expectedIntFieldLabel = "intField"
	intFieldLabel, err = GetTag(test, "IntField", "key")

	assert.Nil(err)
	assert.Equal(expectedIntFieldLabel, intFieldLabel)
}

func TestGetTagMap(t *testing.T) {
	assert := assert.New(t)

	{
		m, err := GetTagMap(test, "label")
		assert.Nil(err)
		assert.Equal(m, expectedMap)
	}
	{
		m, err := GetTagMapReverse(test, "key")
		assert.Nil(err)
		assert.Equal(m, expectedMapReverse)
	}
}
