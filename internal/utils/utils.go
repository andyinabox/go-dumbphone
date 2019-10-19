package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

// CreateTempFile creates a file in the temp dir and returns it
func CreateTempFile(name string) (*os.File, error) {

	// if empty use timestamp
	if name == "" {
		name = Timestamp()
	}

	filename := fmt.Sprintf("%s%s.html", os.TempDir(), name)
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	return f, nil
}

func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// OpenBrowser opens passed url in web browser
func OpenBrowser(url string) error {
	cmd := exec.Command("open", url)

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// BrowserSend open the given file in web browser
func BrowserSend(f *os.File) error {
	return OpenBrowser("file://" + f.Name())
}

// BluetoothSend transfer the given file over bluetooth
func BluetoothSend(f *os.File) error {

	// open in Bluetooth File Exchange
	// http://hints.macworld.com/article.php?story=20040413031046870
	cmd := exec.Command("open", "-a", "/Applications/Utilities/Bluetooth File Exchange.app", f.Name())

	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}

// USBSend transfer the given file over USB
func USBSend(f *os.File) error {
	return errors.New("USBSend is not implemented yet")
}
