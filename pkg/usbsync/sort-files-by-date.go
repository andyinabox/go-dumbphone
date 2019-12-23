package usbsync

import "os"

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
