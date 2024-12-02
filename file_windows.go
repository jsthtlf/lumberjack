package lumberjack

import (
	"os"
	"syscall"
	"time"
)

// createdTime return creation time of the file
func createdTime(info os.FileInfo) time.Time {
	stat, ok := info.Sys().(*syscall.Win32FileAttributeData)
	if !ok {
		return time.Now()
	}

	return time.Unix(0, stat.CreationTime.Nanoseconds())
}
