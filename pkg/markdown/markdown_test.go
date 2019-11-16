package markdown

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	// "io"
	"io/ioutil"
	"os"
	"testing"
)

func TestParse(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	sampleIn, err := os.Open("../../test/data/markdown/sample.md")
	require.Nil(err)
	inBytes, err := ioutil.ReadAll(sampleIn)
	require.Nil(err)

	sampleOut, err := os.Open("../../test/data/markdown/sample.html")
	require.Nil(err)
	outBytes, err := ioutil.ReadAll(sampleOut)
	require.Nil(err)

	parsedBytes, err := Parse(inBytes)
	assert.Nil(err)

	outputPath := "../../test/data/markdown/output.html"
	output, _ := os.Create(outputPath)
	output.Write(parsedBytes)
	defer output.Close()

	if assert.Equal(string(parsedBytes), string(outBytes)) {
		output.Close()
		os.Remove(outputPath)
	}
}
