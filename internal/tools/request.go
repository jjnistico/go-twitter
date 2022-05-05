package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// This function is used to request data from any twitter v2 endpoint.
// In order to use, pass the endpoint (anything after ../v2/ in the api url),
// a method (most like GET) and an optional payload (if applicable)
//
// NOTE: This function returns an array of bytes which you must handle in your
// api endpoint method in order to marshall to JSON
func RequestData(endpoint string, method string, payload io.Reader) ([]byte, error) {
	base_url := os.Getenv("BASE_URL")
	bearer_token := os.Getenv("BEARER_TOKEN")

	req, err := http.NewRequest(method, base_url+endpoint, payload)

	if err != nil {
		fmt.Printf("Error encountered %s %s - %s\n", method, endpoint, err)
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", bearer_token))

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error encountered %s %s - %s\n", method, endpoint, err)
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading response %s\n", err)
		return nil, err
	}

	return body, nil
}
