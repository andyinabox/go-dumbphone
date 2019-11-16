package notes

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

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

func TestParseAll(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	dirname := "../../test/data/notes/"

	notes, err := Create(dirname, 3)
	require.Nil(err)

	files, err := notes.getFiles()
	require.Nil(err)
	require.NotEqual(len(files), 0)

	parsed, err := notes.parseAll(files)
	assert.Nil(err)
	assert.NotEqual(len(parsed), 0, "Parsed files should be greater than 0")

	for _, p := range parsed {
		assert.NotEqual(len(p), 0, "Parsed data should not contain 0 bytes")
	}

}

func TestRun(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	inDir := "../../test/data/notes/"
	outDir := "../../test/data/notes_out/"

	err := os.MkdirAll(outDir, 0777)
	require.Nil(err)

	notes, err := Create(inDir, 3)
	assert.Nil(err)

	md, err := notes.Run()
	assert.Nil(err)

	assert.NotEqual(len(md), 0, "Should return files")
	assert.Equal(len(md), notes.Count, "Should return the expected number of files")
	//
	for _, b := range md {
		assert.NotEqual(len(b), 0, "Bytes read should be greater than zero")
	}
}
