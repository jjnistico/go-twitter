package test

import (
	"fmt"
	"gotwitter/internal/tools/utils/response"
)

func MockResponseFromError(path_param string) []byte {
	response := response.Response{}
	response.AddError(
		"invalid request",
		fmt.Sprintf("`%s` query parameter required", path_param),
		"",
		"response",
	)

	return response.JSON()
}
