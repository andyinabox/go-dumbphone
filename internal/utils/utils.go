package utils

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

// FileSystemDir is an enum for selecting a file system dir
// during USB transfer
type FileSystemDir int

const (
	// DirectionsDir save to directions folder
	DirectionsDir FileSystemDir = iota + 1
	// PodcastDir podcasts folder
	PodcastDir
	// NotesDir notes folder
	NotesDir
	// ReaderDir reader folder
	ReaderDir
)

// USBConnectionError `DUMBP_FS_ROOT` was not found, likely
// that the phone isn't plugged in
type USBConnectionError struct {
}

func (e USBConnectionError) Error() string {
	return "Error connecting to phone via USB. Is it plugged in?"
}

// SortFilesByDate a sort interface ordering a `os.FileInfo` slice
// by date modified
type SortFilesByDate []os.FileInfo

func (f SortFilesByDate) Len() int {
	return len(f)
}
func (f SortFilesByDate) Swap(i, j int) {
	f[i], f[j] = f[j], f[i]
}
func (f SortFilesByDate) Less(i, j int) bool {
	return f[i].ModTime().After(f[j].ModTime())
}

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

// Timestamp creates a Unix-formatted string timestamp
func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}

// OpenBrowser opens passed url in web browser
func OpenBrowser(url string) error {

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
func BrowserSend(f *os.File) error {
	return OpenBrowser("file://" + f.Name())
}

// BluetoothSend transfer the given file over bluetooth
func BluetoothSend(f *os.File) error {

	switch runtime.GOOS {
	case "linux":
		return exec.Command("bluetooth-sendto", f.Name()).Start()
	case "darwin":
		return exec.Command("open", "-a", "/Applications/Utilities/Bluetooth File Exchange.app", f.Name()).Start()
	}

	return errors.New("Unsupported platform")
}

// USBSend transfer the given file over USB
func USBSend(files []os.File, t FileSystemDir) error {
	var savePath string
	switch t {
	case DirectionsDir:
		savePath = os.Getenv("DUMBP_DIRECTIONS_DIR")
	case PodcastDir:
		savePath = os.Getenv("DUMBP_PODCAST_DIR")
	case NotesDir:
		savePath = os.Getenv("DUMBP_NOTES_DIR")
	case ReaderDir:
		savePath = os.Getenv("DUMBP_READER_DIR")
	default:
		savePath = os.Getenv("DUMBP_FS_ROOT")
	}

	if _, err := os.Stat(os.Getenv("DUMBP_FS_ROOT")); os.IsNotExist(err) {
		return USBConnectionError{}
	}

	if _, err := os.Stat(savePath); os.IsNotExist(err) {
		os.MkdirAll(savePath, 0777)
	}

	// for f, _ := range files {

	// }
	return nil
}
