package tools

import (
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
)

type KucoinRequestHeaders struct {
}

func SignRequest(secret string, headers map[string]string, body interface{}, signature_message string) (string, error) {
	hmac_sig := hmac.New(sha256.New, []byte(secret))

	// hmac_sig.Write([]byte(json_body))

	sha := b64.StdEncoding.EncodeToString(hmac_sig.Sum(nil))

	return sha, nil
}
