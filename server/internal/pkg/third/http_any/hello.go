package http_any

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Hello() (map[string]interface{}, error) {
	url := "https://httpbin.org/anything"
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		return nil, err
	}
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	data := make(map[string]interface{})
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}
	return data, nil
}
