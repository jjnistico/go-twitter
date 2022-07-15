package oauth

import "strings"

type RequestTokenResponse struct {
	Token             string
	Secret            string
	CallbackConfirmed bool
}

func parseRequestTokenStringToStruct(str string) RequestTokenResponse {
	strArr := strings.Split(str, "&")
	var tokenResp RequestTokenResponse
	for _, val := range strArr {
		split_val := strings.Split(val, "=")
		switch split_val[0] {
		case "oauth_token":
			tokenResp.Token = split_val[1]
		case "oauth_token_secret":
			tokenResp.Secret = split_val[1]
		case "oauth_callback_confirmed":
			tokenResp.CallbackConfirmed = bool(split_val[1] == "true")
		}
	}
	return tokenResp
}

type AccessTokenResponse struct {
	ScreenName       string `json:"screenName"`
	UserId           string `json:"userId"`
	OAuthToken       string `json:"oauthToken"`
	OAuthTokenSecret string `json:"oauthTokenSecret"`
}

func parseAccessTokenStringToStruct(str string) AccessTokenResponse {
	strArr := strings.Split(str, "&")
	var tokenResp AccessTokenResponse
	for _, val := range strArr {
		split_val := strings.Split(val, "=")
		switch split_val[0] {
		case "oauth_token":
			tokenResp.OAuthToken = split_val[1]
		case "oauth_token_secret":
			tokenResp.OAuthTokenSecret = split_val[1]
		case "user_id":
			tokenResp.UserId = split_val[1]
		case "screen_name":
			tokenResp.ScreenName = split_val[1]
		}
	}
	return tokenResp
}
