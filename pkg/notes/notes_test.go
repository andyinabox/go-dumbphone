package notes

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	// "io"
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

func TestRun(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	var zero int64 = 0

	inDir := "../../test/data/notes/"
	outDir := "../../test/data/notes_out/"

	err := os.MkdirAll(outDir, 0777)
	require.Nil(err)

	notes, err := Create(inDir, 3)
	assert.Nil(err)

	files, err := notes.Run()
	assert.Nil(err)

	assert.NotEqual(len(files), 0, "Should return files")
	assert.Equal(len(files), notes.Count, "Should return the expected number of files")

	for _, src := range files {
		stat, err := src.Stat()
		assert.Nil(err)
		assert.NotEqual(stat.Size(), zero, "Source file size should not be zero")

		// buf := make([]byte, BUFFERSIZE)
		inBytes, err := ioutil.ReadAll(src)
		assert.Nil(err)
		assert.NotEqual(len(inBytes), 0, "Bytes read should be greater than zero")

		newFile := fmt.Sprintf("%s%s", outDir, stat.Name())
		t.Log(newFile)

		err = ioutil.WriteFile(newFile, inBytes, 0777)
		assert.Nil(err)

		// dst, err := os.Create(newFile)
		// assert.Nil(err)
		// defer dst.Close()

		// written, err := io.Copy(dst, src)
		// assert.Nil(err)

		// assert.NotEqual(written, zero, "Bytes written should not be zero")
	}

	for _, f := range files {
		defer f.Close()
		s, err := f.Stat()
		assert.Nil(err)

		outFile := fmt.Sprintf("%s%s", outDir, s.Name())
		outS, err := os.Stat(outFile)
		assert.Nil(err, "File copies should exist")
		assert.NotEqual(outS.Size(), zero, "File size should not be zero")
	}

}
