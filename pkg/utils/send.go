package utils

import (
	"errors"
	"os"

	"github.com/andyinabox/go-dumbphone/pkg/bluetooth"
	"github.com/andyinabox/go-dumbphone/pkg/browser"
	"github.com/andyinabox/go-dumbphone/pkg/usbsync"
)

type SendMethod int

const (
	BLUETOOTH_SEND SendMethod = iota
	USB_SEND
	BROWSER_SEND
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
