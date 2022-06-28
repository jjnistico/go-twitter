package api

import (
	"encoding/json"
	"fmt"
	"gotwitter/internal/endpoint"
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
	// options request used to query available query parameters for endpoint
	// TODO: Move to handler
	if req.Method == http.MethodOptions {
		handleOptionsRequest(w, api_endpoint)
		return
	}

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

func handleOptionsRequest(w http.ResponseWriter, api_endpoint string) {
	options := endpoint.GetEndpointOptions(api_endpoint)

	options_json, err := json.Marshal(
		response.ApiResponseFromData(map[string][]string{"options": options}),
	)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json_resp, err := json.Marshal(response.OneOffErrorResponse(
			err.Error(), "unable to serialize endpoint options",
		))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)

		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(options_json)
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
