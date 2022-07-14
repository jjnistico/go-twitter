package oauth

import (
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
	builder.WriteString("%")

	// hex representation of char as string
	hex_str := fmt.Sprintf("%x", byte(char))
	builder.WriteString(strings.ToUpper(hex_str))
}
