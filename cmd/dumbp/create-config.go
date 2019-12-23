package main

import (
	"io"
	"os"
	"path/filepath"
)

func createConfig(path string, name string, configPath string) (err error) {

	dstFile := filepath.Join(path, name)

	// make config folder
	err = os.MkdirAll(path, 0775)
	if err != nil {
		return
	}

	// open example config
	in, err := os.Open(configPath)
	if err != nil {
		return
	}
	defer in.Close()

	// create config file
	out, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer out.Close()

	// copy file contents
	if _, err = io.Copy(out, in); err != nil {
		return
	}
	err = out.Sync()
	return
}
