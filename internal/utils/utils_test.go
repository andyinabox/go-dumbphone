package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"testing"
	"time"
)

func CreateTempDirWithFiles(name string, n int, sleep int) ([]*os.File, error) {
	dirname := fmt.Sprintf("%s/%s", os.TempDir(), name)

	d, err := time.ParseDuration(strconv.Itoa(sleep) + "ms")
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(dirname, 0777)
	if err != nil {
		return nil, err
	}

	fmt.Println(dirname)

	var files []*os.File

	for i := 0; i < n; i++ {
		filename := fmt.Sprintf("%v/%v-%v.test", dirname, Timestamp(), i)
		fmt.Println(filename)
		file, err := os.Create(filename)
		if err != nil {
			return nil, err
		}
		files = append(files, file)

		if sleep > 0 {
			time.Sleep(d)
		}
	}

	return files, nil
}

func TestTempFile(t *testing.T) {

	f, err := CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating unnamed temp file: %v", err)
	} else {
		t.Logf("Created file %v", f.Name())
	}

	f, err = CreateTempFile("test")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating named temp file: %v", err)
	} else {
		t.Logf("Created file %v", f.Name())
	}

}

func TestBluetoothSend(t *testing.T) {

	if !testing.Verbose() {
		t.SkipNow()
	}

	f, err := CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}

	err = BluetoothSend(f)
	if err != nil {
		t.Errorf("Error sending over BlueTooth: %v", err)
	}
}

func TestSortFilesByDate(t *testing.T) {
	assert := assert.New(t)
	require := require.New(t)

	dirname := Timestamp()
	dirpath := fmt.Sprintf("%s/%s", os.TempDir(), dirname)

	files, err := CreateTempDirWithFiles(dirname, 5, 1000)
	require.Nil(err)

	fi, err := ioutil.ReadDir(dirpath)
	require.Nil(err)

	sort.Sort(SortFilesByDate(fi))

	assert.Equal(len(files), len(fi), "There should be the same number of FileInfo objects as files")

	for i := 1; i < len(fi); i++ {
		mt1 := fi[i-1].ModTime()
		mt2 := fi[i].ModTime()
		// t.Logf("%v, %v", mt1.Unix(), mt2.Unix())
		assert.True(mt2.Before(mt1), "File infos should be in order by time")
	}
}

// func TestBluetoothSend(t *testing.T) {

// }
