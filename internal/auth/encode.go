package auth

import (
	"crypto/hmac"
	"crypto/sha1"
	b64 "encoding/base64"
	"fmt"
	"strings"
)

// algorithm https://developer.twitter.com/en/docs/authentication/oauth-1-0a/percent-encoding-parameters
func percentEncode(str string) string {
	builder := strings.Builder{}
	for _, ch := range str {
		switch ch {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			fallthrough
		case 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z':
			fallthrough
		case 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z':
			fallthrough
		case '-', '.', '_', '~':
			builder.WriteRune(ch)
		default:
			encodeRune(ch, &builder)
		}
	}
	return builder.String()
}

func encodeRune(char rune, builder *strings.Builder) {
	hexStr := fmt.Sprintf("%%%x", byte(char))
	builder.WriteString(strings.ToUpper(hexStr))
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func hmacHash(data string, key string) string {
	mac := hmac.New(sha1.New, []byte(key))
	mac.Write([]byte(data))
	hash := b64.StdEncoding.EncodeToString(mac.Sum(nil))
	return hash
}
