package utils

import (
	"strconv"
	"time"
)

// Timestamp creates a Unix-formatted string timestamp
func Timestamp() string {
	return strconv.FormatInt(time.Now().Unix(), 10)
}
