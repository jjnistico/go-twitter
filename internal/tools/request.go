package tools

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

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
