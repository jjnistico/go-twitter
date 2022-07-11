package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractParameterFromQuery(
	query_params url.Values,
	path_param string) (string, url.Values, error) {

	url_path_parameter := query_params.Get(path_param)

	if len(url_path_parameter) == 0 {
		return "", nil, fmt.Errorf("`%s` query parameter required", path_param)
	}

	query_params.Del(path_param)

	return url_path_parameter, query_params, nil
}

func VerifyRequiredQueryParams(query_params url.Values, required_params []string) error {
	if len(required_params) == 0 {
		return nil
	}

	missing_query_params := []string{}
	for _, required_param := range required_params {
		curr_query_val := query_params[required_param]
		if len(curr_query_val) == 0 {
			missing_query_params = append(missing_query_params, required_param)
		}
	}

	if len(missing_query_params) > 0 {
		return fmt.Errorf("missing query parameters: [%s]", strings.Join(missing_query_params, ", "))
	}

	return nil
}
