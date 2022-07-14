package types

type createResponseData interface{ CreateTweetResponse }

type deleteResponseData interface{ DeleteTweetResponse }

type getResponseData interface {
	TweetsResponse | UsersResponse
}

type ResponseData interface {
	createResponseData | deleteResponseData | getResponseData
}
