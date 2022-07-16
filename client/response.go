package gotwit

type createResponseData interface{ createTweetResponse }

type deleteResponseData interface{ deleteTweetResponse }

type getResponseData interface {
	tweetsResponse | usersResponse | userResponse
}

type responseData interface {
	createResponseData | deleteResponseData | getResponseData
}
