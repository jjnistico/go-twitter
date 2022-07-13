package api

import (
	"encoding/json"
	"gotwitter/internal/server"
	"io"
	"net/url"
)

func ApiRequest[T server.ResponseT](
	endpoint string,
	method string,
	query_params url.Values,
	required_params []string,
	payload io.Reader,
) *server.GOTResponse[T] {
	request := server.NewRequest(endpoint, query_params, method, payload)

	request.VerifyQueryParams(required_params)

	request.Authorize()

	data, errors, status_code := request.Execute()

	var um_data T
	if err := json.Unmarshal(data.([]byte), &um_data); err != nil {
		panic(err)
	}

	return server.NewResponse(um_data, errors, status_code)
}
