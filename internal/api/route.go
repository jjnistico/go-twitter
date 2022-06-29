package api

import (
	"encoding/json"
	"fmt"
	"gotwitter/internal/tools"
	"gotwitter/internal/tools/utils/response"
	"net/http"
	"strings"
)

func ApiRoute(
	w http.ResponseWriter,
	req *http.Request,
	api_endpoint string,
	http_method string,
	required_query_params []string) {

	if len(required_query_params) > 0 {
		has_required_params := hasRequiredQueryParams(w, req, required_query_params)

		if !has_required_params {
			return
		}
	}

	data, status_code, err := tools.RequestData(api_endpoint, req.URL.Query(), http_method, req.Body)

	w.WriteHeader(status_code)

	if err != nil {
		json_resp, err := json.Marshal(response.OneOffErrorResponse(err.Error(), "error requesting data"))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)

		return
	}

	w.Write(data)
}

func hasRequiredQueryParams(w http.ResponseWriter, req *http.Request, required_params []string) bool {
	missing_query_params := []string{}
	for _, required_param := range required_params {
		curr_query_val := req.URL.Query().Get(required_param)
		if len(curr_query_val) == 0 {
			missing_query_params = append(missing_query_params, required_param)
		}
	}

	if len(missing_query_params) > 0 {
		w.WriteHeader(http.StatusBadRequest)

		json_resp, err := json.Marshal(response.ErrorResponse(
			fmt.Sprintf("missing parameters: [%s]", strings.Join(missing_query_params, ", ")),
			"invalid request",
			"missing required query parameter",
			req.Method,
		))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)
		return false
	}
	return true
}
