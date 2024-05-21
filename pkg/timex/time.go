package timex

import "time"

func CurDate() string {
	return time.Now().Format("2006-01-02")
}

func CurDatetime() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
