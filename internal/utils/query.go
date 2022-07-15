package utils

import (
	"fmt"
	"net/url"
	"strings"
)

func ExtractParameterFromQuery(queryParams url.Values, pathParam string) (string, url.Values, error) {
	urlPathParameter := queryParams.Get(pathParam)

	if len(urlPathParameter) == 0 {
		return "", nil, fmt.Errorf("`%s` query parameter required", pathParam)
	}

	queryParams.Del(pathParam)

	return urlPathParameter, queryParams, nil
}

func VerifyRequiredQueryParams(queryParams url.Values, requiredParams []string) error {
	if len(requiredParams) == 0 {
		return nil
	}

	missingQueryParams := []string{}
	for _, requiredParam := range requiredParams {
		currQueryVal := queryParams[requiredParam]
		if len(currQueryVal) == 0 {
			missingQueryParams = append(missingQueryParams, requiredParam)
		}
	}

	if len(missingQueryParams) > 0 {
		return fmt.Errorf("missing query parameters: [%s]", strings.Join(missingQueryParams, ", "))
	}

	return nil
}
