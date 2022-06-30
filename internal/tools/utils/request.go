package utils

import (
	"fmt"
	"gotwitter/internal/tools/utils/response"
	"net/http"
)

func GetPathParameterFromQuery(w http.ResponseWriter, req *http.Request, path_param string) string {
	url_path_parameter := req.URL.Query().Get(path_param)

	if len(url_path_parameter) == 0 {
		w.WriteHeader(http.StatusBadRequest)

		response := response.Response{}
		response.AddError("invalid request",
			fmt.Sprintf("`%s` query parameter required", path_param), "", "request",
		)

		w.Write(response.JSON())
		return ""
	}

	query_params := req.URL.Query()
	query_params.Del(path_param)
	req.URL.RawQuery = query_params.Encode()

	return url_path_parameter
}
