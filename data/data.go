// Code generated for package data by go-bindata DO NOT EDIT. (@generated)
// sources:
// bin/data/directions.html
// bin/data/reader.html
package data

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func bindataRead(data []byte, name string) ([]byte, error) {
	gz, err := gzip.NewReader(bytes.NewBuffer(data))
	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}

	var buf bytes.Buffer
	_, err = io.Copy(&buf, gz)
	clErr := gz.Close()

	if err != nil {
		return nil, fmt.Errorf("Read %q: %v", name, err)
	}
	if clErr != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _binDataDirectionsHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x9c\x55\xd1\x72\xab\x36\x10\x7d\xe7\x2b\xb6\xee\x4b\xef\xcc\xb5\xf1\xcd\xa4\x79\x20\x0a\x33\x69\xd3\x4e\x3a\x93\xa4\x9d\xda\x3f\x20\x90\x02\x9a\x08\x89\x4a\x8b\x63\x0f\xe5\xdf\x3b\x42\x08\xb0\xe3\xa4\x69\x9f\x58\x56\xbb\x67\xcf\xee\x1e\x01\xf9\x8e\xe9\x1c\x0f\x35\x87\x12\x2b\x99\x46\xc4\x3f\x22\x52\x72\xca\xd2\x08\x80\xa0\x40\xc9\xd3\x3b\x61\x78\x8e\x42\x2b\x4b\x62\xef\x71\x67\x16\x0f\xde\x02\x10\x55\x01\x6d\x6f\x01\x30\x61\x6b\x49\x0f\x09\x64\x52\xe7\x2f\xd7\x83\xb7\xa2\xa6\x10\x2a\x81\x6f\xeb\x7a\x0f\x6b\xef\xed\xa2\xfe\xa1\xe5\x98\x1b\xa2\xd6\x21\xad\xa6\x8c\x09\x55\xcc\x3c\x52\x58\x5c\xf6\xa5\x97\xb5\xb6\xc2\xb1\x4a\x40\x28\x2b\x18\x3f\x45\xd5\xf2\x6b\x30\x9b\xa9\xc6\x00\xb9\x94\xfc\x19\x13\xf8\xb1\xde\x1f\xa5\x61\x39\x06\x22\xdf\xe3\x92\x4a\x51\xa8\x04\x5c\x70\x60\xb0\xe3\x06\x45\x4e\x65\x38\x43\x5d\x1f\x43\x98\x11\x22\xd3\x86\x71\xb3\xcc\x34\xa2\xae\x12\xf8\x56\xef\xc1\x6a\x29\x18\x7c\xcf\x18\x3b\x4e\xa2\x99\xe4\x27\x73\x98\xf2\xd6\x13\x4b\x00\x12\x0f\x93\x27\xb1\xdf\x53\x44\x32\xcd\x0e\xa9\x03\x22\x6e\x13\xd6\xe4\x37\x0b\x46\x91\x26\xa2\xa2\x05\x8f\x6b\x55\x5c\x67\xd4\xf2\xab\xcb\xaf\x6d\x0b\xab\xdf\x77\xdc\xec\x04\x7f\xfd\x43\xcb\x83\x14\x8a\xc3\xdf\x50\x0f\xe6\x56\xff\x74\x75\x09\x5d\xb7\x80\x1e\xad\x6d\x0d\x55\x05\x87\xd5\x03\x2f\x6c\xd7\x33\x25\x3d\x51\xbf\x75\x82\xbe\xae\xe7\x4c\xd0\x04\xd3\xbd\x94\xe9\xaf\x46\x57\x24\xc6\x72\xee\x65\xa9\x63\xb0\x41\x6a\xf0\x96\x31\xc3\xad\x85\xae\x23\x31\xb2\x11\x25\x9e\x60\xde\x20\x6e\xf5\x79\xbc\x5f\x14\xfb\x1f\x68\x77\xc2\x22\x55\x39\x3f\x8f\x19\x4e\x57\xf7\x4d\x45\xd5\x9f\x9c\xb2\x7e\x43\xff\x85\xad\xa8\x84\x2a\xde\x41\x6f\x0c\x75\xd2\x85\xae\x1b\xcf\xdc\xbc\xc5\x33\xf0\xbf\x00\x0d\x55\x56\xe0\xa3\x66\x1c\x16\xc3\xcb\xe2\x38\x14\xe0\x87\x1e\x87\xd7\xd4\x60\x63\xf8\x56\x54\x6e\x91\xa5\x23\xdb\xdb\x5d\x07\x4b\x68\xdb\xd5\xad\x31\x62\x47\xe5\xe9\x79\xd7\x7d\x39\x2a\xcc\x15\x9b\xe1\x9f\xef\x91\xc4\xe3\xc2\x49\x3c\xe8\xa0\xd7\x44\x79\x91\x6e\x9a\xaa\xa2\xe6\x40\xe2\xf2\xc2\x0b\x51\x4b\x9f\x34\x6a\x68\x83\xbc\xb6\x43\x0d\x22\x45\x3f\x86\xfb\xed\xe3\xc3\x6f\xca\xa2\x69\xfc\xe7\xa5\x1f\xaf\x14\x21\x33\x90\x22\xb1\x47\x8b\x86\x11\x31\x8e\x54\x48\xce\xa6\xef\xd2\x20\xce\xd2\xa4\x81\xd0\xdd\x10\x03\xf3\x8f\x57\x20\xd7\x7c\x48\x2e\x0b\xfd\x7e\xe2\x32\x7d\x74\x89\xd2\x13\xb8\xd0\x98\xab\xfb\x2a\xb0\x84\xd5\xd6\xef\xd6\x73\xb5\xd3\x86\x49\x06\xfd\x1d\xbf\x59\xe4\x5a\x6a\x93\x80\x2b\xf5\x20\x14\x5f\x6d\xf9\x1e\x7f\x76\x3e\xe8\xba\x6b\xc8\x68\xfe\x52\x18\xdd\x28\x36\x0b\x09\xc7\x8b\x74\x74\x6d\x4a\x6d\xf0\x89\x56\x5e\xbf\xd9\x44\xc3\x4f\x78\x7c\x3d\xbb\x91\xe8\xb3\xa4\xc3\x54\x43\xb3\x24\x4b\x8f\x34\xba\x41\x5d\xaf\x66\x2c\xfe\x5d\xc2\x5f\x26\x35\xcc\x65\xf3\xd4\x54\x0e\xcb\xd5\x06\xeb\x8c\xb7\x61\x43\xed\x41\xfd\xe7\x2b\xbf\x73\x35\x4e\xeb\x92\x78\x6a\xac\x6d\x81\x2b\x36\xeb\x59\xcf\x7a\x3e\xab\xa6\x4f\xcb\xfd\xed\x3d\x0c\xb2\x77\xd6\xb9\x4b\xe1\x68\x45\x47\x5b\x1c\xcd\x88\xc4\xc3\x3f\x81\xc4\xfd\x4f\xfd\x9f\x00\x00\x00\xff\xff\xe8\x86\xc3\x62\xeb\x07\x00\x00")

func binDataDirectionsHtmlBytes() ([]byte, error) {
	return bindataRead(
		_binDataDirectionsHtml,
		"bin/data/directions.html",
	)
}

func binDataDirectionsHtml() (*asset, error) {
	bytes, err := binDataDirectionsHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bin/data/directions.html", size: 2027, mode: os.FileMode(420), modTime: time.Unix(1581808929, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _binDataReaderHtml = []byte("\x1f\x8b\x08\x00\x00\x00\x00\x00\x00\xff\x5c\x52\xcd\x4e\x84\x30\x10\xbe\xf7\x29\xc6\x78\x15\x81\x83\x17\x9c\x70\xf0\xec\xc9\xe8\x03\x14\xa6\xc2\xc4\x42\x1b\x3a\xbb\x81\xac\xfb\xee\x86\x16\x76\xdd\x3d\xf5\xcb\xf7\x37\x4d\xa7\xf8\x40\xae\x95\xc5\x1b\xe8\x65\xb0\xb5\xc2\x74\x28\xec\x8d\xa6\x5a\x01\xa0\xb0\x58\x53\x9f\x4e\xf0\xfc\xb9\x22\x38\x9f\x31\x4f\xdc\xaa\x06\x59\x12\x02\xe0\xa1\x83\x53\x44\x00\xc4\xc1\x5b\xbd\x54\xd0\x58\xd7\xfe\xbc\x6e\xec\xa0\xa7\x8e\xc7\x0a\xca\xc2\xcf\x50\x24\xf6\xac\xe2\xe1\xec\x25\xbb\xbb\x8a\x3d\xe6\x35\x11\x8f\xdd\x3f\xc6\x72\x90\x2c\x8e\xce\xbc\x0b\x2c\xec\xc6\x0a\x78\x0c\x4c\xe6\xbe\xd5\xd9\xa7\x1d\x1e\xae\x33\xb6\xca\xcc\x9a\x6f\xa9\xe0\xc5\xcf\x37\x31\xe9\x2f\x46\x31\xb3\x64\xda\x72\x37\x56\xb0\x9a\xf7\x1b\x1c\xcd\x24\xdc\x6a\xbb\x6b\xe2\xfc\x6d\xc5\x74\xa9\x68\xdc\x44\x66\xca\x1a\x27\xe2\x86\x0a\x4a\x3f\x43\x70\x96\x09\x1e\x89\xe8\x36\xa4\x1b\x6b\xee\xde\xe1\x9a\x2b\xae\xb7\x04\xc0\x7c\x7b\x79\xcc\xd3\xa6\x14\x36\x8e\x96\x3a\x15\x61\x5f\xde\x6d\xac\x2f\xd3\x92\xd0\xd7\xc8\x51\xfb\xfa\x78\x8f\x0a\xd7\x98\xfb\x24\x26\x07\xf1\x31\x1a\xde\x1c\x2d\xf0\x0b\x87\xd1\x84\x56\xfb\x54\xb3\x6a\x4a\x29\xcc\xb7\x61\x98\xc7\xff\xf2\x17\x00\x00\xff\xff\x9e\x9a\x9a\x8d\x46\x02\x00\x00")

func binDataReaderHtmlBytes() ([]byte, error) {
	return bindataRead(
		_binDataReaderHtml,
		"bin/data/reader.html",
	)
}

func binDataReaderHtml() (*asset, error) {
	bytes, err := binDataReaderHtmlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "bin/data/reader.html", size: 582, mode: os.FileMode(420), modTime: time.Unix(1581808942, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"bin/data/directions.html": binDataDirectionsHtml,
	"bin/data/reader.html":     binDataReaderHtml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"bin": &bintree{nil, map[string]*bintree{
		"data": &bintree{nil, map[string]*bintree{
			"directions.html": &bintree{binDataDirectionsHtml, map[string]*bintree{}},
			"reader.html":     &bintree{binDataReaderHtml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
