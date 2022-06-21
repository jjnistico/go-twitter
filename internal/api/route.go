package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotwitter/internal/endpoint"
	"gotwitter/internal/tools"
	"gotwitter/internal/tools/utils"
	"net/http"
)

func ApiRoute(
	w http.ResponseWriter,
	req *http.Request,
	api_endpoint string,
	http_method string,
	payload interface{},
	required_query_params []string) {
	// options request used to query available query parameters for endpoint
	if req.Method == http.MethodOptions {
		handleOptionsRequest(w, api_endpoint)
		return
	}

	if len(required_query_params) > 0 {
		missing_query_params := []string{}
		for _, required_param := range required_query_params {
			curr_query_val := req.URL.Query().Get(required_param)
			if len(curr_query_val) == 0 {
				missing_query_params = append(missing_query_params, required_param)
			}
		}

		if len(missing_query_params) > 0 {
			w.WriteHeader(http.StatusBadRequest)

			json_resp, err := json.Marshal(utils.ErrorResponse(missing_query_params, "missing required query parameter"))

			if err != nil {
				panic(err)
			}

			w.Write(json_resp)
			return
		}
	}

	var buf bytes.Buffer
	if payload == nil {
		payload = map[string]string{}
	}
	err := json.NewEncoder(&buf).Encode(payload)

	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)

		json_resp, err := json.Marshal(utils.OneOffErrorResponse(err.Error(), "unable to buffer request payload"))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)
		return
	}

	data, status_code, err := tools.RequestData(api_endpoint, req.URL.Query(), http_method, &buf)

	w.WriteHeader(status_code)

	if err != nil {
		fmt.Println(err)

		json_resp, err := json.Marshal(utils.OneOffErrorResponse(err.Error(), "error requesting data"))

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

	options_map := make(map[string][]string)
	options_map["options"] = options
	options_json, err := json.Marshal(utils.ApiResponseFromData(options_map))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)

		json_resp, err := json.Marshal(utils.OneOffErrorResponse(err.Error(), "unable to serialize endpoint options"))

		if err != nil {
			panic(err)
		}

		w.Write(json_resp)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(options_json)
}
