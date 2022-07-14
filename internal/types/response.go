package types

type createResponseData interface{ CreateTweetResponse }

type deleteResponseData interface{ DeleteTweetResponse }

type getResponseData interface {
	TweetsResponse | UsersResponse | UserResponse
}

type ResponseData interface {
	createResponseData | deleteResponseData | getResponseData
}
