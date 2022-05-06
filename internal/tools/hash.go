package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
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
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func HmacHash(data string, key string) string {
	mac := hmac.New(sha256.New, []byte(key))
	mac.Write([]byte(data))
	hash := hex.EncodeToString(mac.Sum(nil))
	return hash
}
