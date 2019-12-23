package transfer

import (
	"errors"
	"os"

	"github.com/andyinabox/go-dumbphone/pkg/bluetooth"
	"github.com/andyinabox/go-dumbphone/pkg/browser"
	"github.com/andyinabox/go-dumbphone/pkg/usbsync"
)

func Send(f *os.File, method SendMethod, path string) error {
	switch method {
	case BLUETOOTH_SEND:
		return bluetooth.Send(f)
	case USB_SEND:
		return usbsync.Send(f, path)
	case BROWSER_SEND:
		return browser.OpenFile(f)
	}
	return errors.New("Inavalid send method")
}
