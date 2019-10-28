package notes

import (
	// "bytes"
	"errors"
	"fmt"
	"github.com/andyinabox/go-dumbphone/internal/utils"
	// "github.com/microcosm-cc/bluemonday"
	"gopkg.in/russross/blackfriday.v2"
	// "io"
	"io/ioutil"
	"os"
	"sort"
)

// Collection Struct for notes sync
type Collection struct {
	SourceDir string
	Count     int
}

// Create a new notes collection
func Create(dir string, count int) (*Collection, error) {

	// validate directory
	fi, err := os.Stat(dir)
	if err != nil {
		return nil, err
	}
	if !fi.IsDir() {
		return nil, errors.New("Path is not a directory")
	}
	if count < 1 {
		return nil, errors.New("Count should be 1 or more")
	}

	collection := &Collection{
		SourceDir: dir,
		Count:     count,
	}

	return collection, nil
}

// Parse write html to io writer from markdown
func Parse(b []byte) ([]byte, error) {
	output := blackfriday.Run(b)
	// html := bluemonday.UGCPolicy().SanitizeBytes(output)
	return output, nil
}

// Run will parsee all files into temp files, then return
// the file list
func (c *Collection) Run() ([]*os.File, error) {

	files, err := c.getFiles()
	if err != nil {
		return nil, err
	}

	parsed, err := c.parseAll(files)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}

func (c *Collection) getFiles() ([]*os.File, error) {
	// make sure containing dir exists
	if _, err := os.Stat(c.SourceDir); os.IsNotExist(err) {
		return nil, err
	}

	// collect file info for all files
	fi, err := ioutil.ReadDir(c.SourceDir)
	if err != nil {
		return nil, err
	}

	// sort by modified date
	sort.Sort(utils.SortFilesByDate(fi))

	var files []*os.File

	i := 0
	for _, f := range fi {
		// break loop if we've reached our limit
		if i >= c.Count {
			break
		}
		// skip to the next iteration if a directory
		if f.IsDir() {
			continue
		}

		// open file and add to files list
		filepath := fmt.Sprintf("%v/%v", c.SourceDir, f.Name())
		f, err := os.Open(filepath)
		if err != nil {
			return nil, err
		}
		files = append(files, f)
		i++
	}

	return files, nil
}

// func byDate()

// func (c *Collection) verifyRequirements() error {
// 	if c.SourceDir == "" || c.Count == nil {
// 		return errors.New("Missing required attributes")
// 	}
// }

func (c *Collection) parseAll(files []*os.File) ([]*os.File, error) {

	var parsed []*os.File

	for _, f := range files {

		// get file info
		s, err := f.Stat()
		if err != nil {
			return nil, err
		}

		// read file into bytes
		b, err := ioutil.ReadAll(f)
		if err != nil {
			return nil, err
		}

		// create temp file to store output
		tmp, err := utils.CreateTempFile(s.Name())
		if err != nil {
			return nil, err
		}

		// parse markdown into html bytes
		p, err := Parse(b)
		if err != nil {
			return nil, err
		}

		// write parsed bytes to temp file
		_, err = tmp.Write(p)
		if err != nil {
			return nil, err
		}

		// add to parsed files slice
		parsed = append(parsed, tmp)
	}

	return parsed, nil
}
