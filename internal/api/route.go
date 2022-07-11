package api

import (
	"gotwitter/internal/server"
	"io"
	"net/url"
)

func ApiRequest(
	endpoint string,
	method string,
	query_params url.Values,
	required_params []string,
	payload io.Reader,
) *server.GOTResponse {
	request := server.NewRequest(endpoint, query_params, method, payload)

	request.VerifyQueryParams(required_params)

	request.Authorize()

	data, status, errors := request.Execute()

	response := server.NewResponse(data, errors, status)

	return response
}
