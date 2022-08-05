package auth

import "time"

type tsFunc func() int64

var unixFunc tsFunc = func() int64 {
	return time.Now().Unix()
}

func resetClockImpl() {
	unixFunc = func() int64 {
		return time.Now().Unix()
	}
}

func unixTimestamp() int64 {
	return unixFunc()
}
