package lumberjack

import (
	"os"
	"syscall"
	"time"
)

// createdTime return creation time of the file
func createdTime(info os.FileInfo) time.Time {
	stat, ok := info.Sys().(*syscall.Stat_t)
	if !ok {
		return time.Now()
	}

	return time.Unix(stat.Atim.Sec, stat.Atim.Nsec)
}
