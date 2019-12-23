package utils

import (
	"fmt"
	"os"
)

// CreateTempFile creates a file in the temp dir and returns it
func CreateTempFile(name string) (*os.File, error) {

	// if empty use timestamp
	if name == "" {
		name = Timestamp()
	}

	filename := fmt.Sprintf("%s/%s.html", os.TempDir(), name)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}
