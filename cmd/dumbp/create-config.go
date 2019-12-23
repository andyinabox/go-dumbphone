package main

import (
	"io"
	"os"
	"path/filepath"
)

const ExampleConfig string = "../../config.example.yml"

func createConfig(path string, name string) (err error) {

	dstFile := filepath.Join(path, name)

	// make config folder
	err = os.MkdirAll(path, 0700)
	if err != nil {
		return
	}

	// open example config
	in, err := os.Open(ExampleConfig)
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
