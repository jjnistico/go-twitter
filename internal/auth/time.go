package auth

import "time"

type tsFunc func() int64

var unixFunc tsFunc

func init() {
	resetClockImpl()
}

func resetClockImpl() {
	unixFunc = func() int64 {
		return time.Now().Unix()
	}
}

func unixTime() int64 {
	return unixFunc()
}
