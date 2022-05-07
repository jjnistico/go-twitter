package tools

import (
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"math/rand"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func GenerateNonce(length int) string {
	b := make([]byte, length)
	for i := 0; i < length; {
		if idx := int(rand.Int31() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func HmacHash(data string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	hash := b64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hash
}
