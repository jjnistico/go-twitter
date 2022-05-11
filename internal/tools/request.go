package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func RequestData(api_url string, query_params string, method string, payload io.Reader) ([]byte, error) {
	bearer_token := os.Getenv("BEARER_TOKEN")

	req, err := http.NewRequest(method, api_url+"?"+url.QueryEscape(query_params), payload)

	if err != nil {
		return nil, fmt.Errorf("error generating request for %s", err.Error())
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_token))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, fmt.Errorf("error calling %s", err.Error())
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error reading response from %s", api_url)
	}

	return data, nil
}

func BuildAuthorizationHeader(header_entries []map[string]string) string {
	// header string is fields joined by ", ". All key/values are percent encoded
	builder := strings.Builder{}
	builder.WriteString("OAuth ")
	for idx, entry := range header_entries {
		for k, v := range entry {
			if v == "" {
				continue
			}
			builder.WriteString(fmt.Sprintf("%s=\"%s\"", PercentEncode(k), PercentEncode(v)))
		}
		if idx < len(header_entries)-1 {
			builder.WriteString(", ")
		}
	}

	return builder.String()
}
