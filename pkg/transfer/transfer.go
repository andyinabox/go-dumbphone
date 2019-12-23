package transfer

type SendMethod int

const (
	BLUETOOTH_SEND SendMethod = iota
	USB_SEND
	BROWSER_SEND
)
