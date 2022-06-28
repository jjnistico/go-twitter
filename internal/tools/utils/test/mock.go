package test

import (
	"encoding/json"
	"fmt"
	"gotwitter/internal/tools/utils/response"
)

func MockResponseFromError(path_param string) []byte {
	json_resp, _ := json.Marshal(response.OneOffErrorResponse(
		fmt.Sprintf("`%s` query parameter required", path_param), "invalid request",
	))

	return json_resp
}
