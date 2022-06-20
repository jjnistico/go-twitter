package tools

import (
	"fmt"
	"gotwitter/internal/tools/oauth"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func RequestData(
	api_url string,
	query_params url.Values,
	method string,
	payload io.Reader,
) ([]byte, int, error) {
	req, err := http.NewRequest(method, api_url+"?"+query_params.Encode(), payload)
	req.Header.Add("content-type", "application/json")

	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf(
			"error generating request for api: %s | query: %s -> %s",
			api_url,
			query_params.Encode(),
			err.Error(),
		)
	}

	oauth.AuthorizeRequest(req)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, resp.StatusCode, fmt.Errorf("error calling %s", err.Error())
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, http.StatusInternalServerError, fmt.Errorf("error reading response from %s", api_url)
	}

	return data, resp.StatusCode, nil
}

type EMPTY_PAYLOAD = map[string]string
