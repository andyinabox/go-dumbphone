package bluetooth

import (
	"errors"
	"os"
	"os/exec"
	"runtime"
)

func Send(f *os.File) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("bluetooth-sendto", f.Name()).Start()
	case "darwin":
		return exec.Command("open", "-a", "/Applications/Utilities/Bluetooth File Exchange.app", f.Name()).Start()
	}

	return errors.New("Unsupported platform")
}
