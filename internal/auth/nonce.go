package auth

import "math/rand"

type genNonce func(len int) string

var nonceFunc genNonce

func init() {
	resetNonceImpl()
}

func resetNonceImpl() {
	nonceFunc = func(len int) string {
		return generateNonce(len)
	}
}

func getNonce(len int) string {
	return nonceFunc(len)
}

func generateNonce(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; {
		if idx := int(rand.Int31() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}
