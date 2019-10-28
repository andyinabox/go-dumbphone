package notes

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCreate(t *testing.T) {
	assert := assert.New(t)
	dirname := "../../test/data/notes/"

	_, err := Create("./THIS_DOES_NOT_EXIST", 3)
	assert.NotNil(err, "Should throw an error for a file that doesn't exist")

	_, err = Create("./notes.go", 3)
	assert.NotNil(err, "Should throw an error for a file that exists but isn't a dir")

	_, err = Create("dirname", 0)
	assert.NotNil(err, "Should throw an error if count is less than 1")

	_, err = Create(dirname, 3)
	assert.Nil(err, "Should not throw error for dir that exists and count greater than 0")

}

func TestSortFiles(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	d, err := time.ParseDuration("1s")
	require.Nil(err)

	dirname := "../../test/data/notes/"

	matches, err := filepath.Glob(dirname + "*")

	for _, m := range matches[:5] {
		err = os.Chtimes(m, time.Now(), time.Now())
		require.Nil(err)
		time.Sleep(d)
	}

	notes, err := Create(dirname, 3)
	require.Nil(err)

	files, err := notes.getFiles()
	assert.Nil(err)

	for _, f := range files {
		stat, err := f.Stat()
		assert.Nil(err)
		t.Logf("%v: %v\n", stat.Name(), stat.ModTime().Unix())
	}

	assert.Equal(len(files), notes.Count, "The returned files should match the desired count")

}

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
