package api

import (
	"encoding/json"
	"gotwitter/internal/network"
	"gotwitter/internal/types"
	"io"
	"net/url"
	"strings"
)

func ApiRequest[T network.ResponseData](
	endpoint string,
	method string,
	options types.GOTOptions,
	required_params []string,
	payload io.Reader,
) *network.GOTResponse[T] {
	query_params := url.Values{}

	for k, v := range options {
		query_params.Set(k, strings.Join(v, ","))
	}

	request := network.NewRequest(endpoint, query_params, method, payload)

	request.VerifyQueryParams(required_params)

	request.Authorize()

	data, errors := request.Execute()

	var structured_data T

	if len(errors) == 0 {
		if err := json.Unmarshal(data.([]byte), &structured_data); err != nil {
			panic(err.Error())
		}
	}

	return network.NewResponse(structured_data, errors)
}
