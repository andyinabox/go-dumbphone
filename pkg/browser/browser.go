package browser

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

// OpenBrowser opens passed url in web browser
func Open(url string) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	}

	return errors.New("Unsupported platform")
}

// BrowserSend open the given file in web browser
func OpenFile(f *os.File) error {
	return Open("file://" + f.Name())
}
