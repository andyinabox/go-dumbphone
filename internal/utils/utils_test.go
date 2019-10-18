package utils

import (
	"testing"
)

func TestTempFile(t *testing.T) {

	f, err := CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating unnamed temp file: %v", err)
	} else {
		t.Logf("Created file %v", f.Name())
	}

	f, err = CreateTempFile("test")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating named temp file: %v", err)
	} else {
		t.Logf("Created file %v", f.Name())
	}

}

func TestBluetoothSend(t *testing.T) {

	if !testing.Verbose() {
		t.SkipNow()
	}

	f, err := CreateTempFile("")
	defer f.Close()
	if err != nil {
		t.Errorf("Error creating file: %v", err)
	}

	err = BluetoothSend(f)
	if err != nil {
		t.Errorf("Error sending over BlueTooth: %v", err)
	}
}

// func TestBluetoothSend(t *testing.T) {

// }
