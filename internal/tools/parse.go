package tools

import "strings"

type TokenResponse struct {
	Token             string
	Secret            string
	CallbackConfirmed bool
}

func ParseTokenStringToStruct(str string) TokenResponse {
	str_arr := strings.Split(str, "&")
	var token_resp TokenResponse
	for _, val := range str_arr {
		split_val := strings.Split(val, "=")
		switch split_val[0] {
		case "oauth_token":
			token_resp.Token = split_val[1]
		case "oauth_token_secret":
			token_resp.Secret = split_val[1]
		case "oauth_callback_confirmed":
			token_resp.CallbackConfirmed = bool(split_val[1] == "true")
		}
	}
	return token_resp
}
