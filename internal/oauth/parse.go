package oauth

import "strings"

type RequestTokenResponse struct {
	Token             string
	Secret            string
	CallbackConfirmed bool
}

func ParseRequestTokenStringToStruct(str string) RequestTokenResponse {
	str_arr := strings.Split(str, "&")
	var token_resp RequestTokenResponse
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

type AccessTokenResponse struct {
	ScreenName       string `json:"screenName"`
	UserId           string `json:"userId"`
	OAuthToken       string `json:"oauthToken"`
	OAuthTokenSecret string `json:"oauthTokenSecret"`
}

func ParseAccessTokenStringToStruct(str string) AccessTokenResponse {
	str_arr := strings.Split(str, "&")
	var token_resp AccessTokenResponse
	for _, val := range str_arr {
		split_val := strings.Split(val, "=")
		switch split_val[0] {
		case "oauth_token":
			token_resp.OAuthToken = split_val[1]
		case "oauth_token_secret":
			token_resp.OAuthTokenSecret = split_val[1]
		case "user_id":
			token_resp.UserId = split_val[1]
		case "screen_name":
			token_resp.ScreenName = split_val[1]
		}
	}
	return token_resp
}
