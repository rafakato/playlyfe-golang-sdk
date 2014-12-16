package playlyfe

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const apiBaseUrl string = "https://api.playlyfe.com/v1"

type PlaylyfeClient struct {
	http http.Client
}

type Endpoint struct {
	Url             string
	QueryParameters map[string]string
	RequestBody     interface{}
}

func newEndpointRequest(verb string, endpoint Endpoint) (req *http.Request, err error) {
	if verb != "GET" && verb != "POST" && verb != "PUT" && verb != "DELETE" {
		err = errors.New("Invalid verb")
		return
	}

	var urlParts []string
	urlParts = append(urlParts, apiBaseUrl)
	urlParts = append(urlParts, endpoint.Url)

	if endpoint.QueryParameters != nil {
		urlParts = append(urlParts, "?")
		for k, v := range endpoint.QueryParameters {
			urlParts = append(urlParts, k)
			urlParts = append(urlParts, "=")
			urlParts = append(urlParts, v)
		}
	}
	endpointUrl := strings.Join(urlParts, "")

	var jsonBody []byte
	if endpoint.RequestBody != nil {
		jsonBody, err = json.Marshal(endpoint.RequestBody)
	} else {
		jsonBody = nil
	}

	req, err = http.NewRequest(verb, endpointUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		return
	}
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func (client *PlaylyfeClient) GetRaw(endpoint Endpoint) (result string, err error) {
	req, err := newEndpointRequest("GET", endpoint)
	if err != nil {
		return
	}

	resp, err := client.http.Do(req)
	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

func (client *PlaylyfeClient) Get(endpoint Endpoint, dataStructResult interface{}) (err error) {
	body, err := client.GetRaw(endpoint)

	err = json.Unmarshal([]byte(body), &dataStructResult)

	return
}

func (client *PlaylyfeClient) PostRaw(endpoint Endpoint) (result string, err error) {
	req, err := newEndpointRequest("POST", endpoint)
	if err != nil {
		return
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

func (client *PlaylyfeClient) Post(endpoint Endpoint, dataStructResult interface{}) (err error) {
	body, err := client.PostRaw(endpoint)

	err = json.Unmarshal([]byte(body), &dataStructResult)

	return
}

func (client *PlaylyfeClient) PutRaw(endpoint Endpoint) (result string, err error) {
	req, err := newEndpointRequest("PUT", endpoint)
	if err != nil {
		return
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

func (client *PlaylyfeClient) Put(endpoint Endpoint, dataStructResult interface{}) (err error) {
	body, err := client.PutRaw(endpoint)

	err = json.Unmarshal([]byte(body), &dataStructResult)

	return
}

func (client *PlaylyfeClient) DeleteRaw(endpoint Endpoint) (result string, err error) {
	req, err := newEndpointRequest("DELETE", endpoint)
	if err != nil {
		return
	}

	resp, err := client.http.Do(req)

	if err != nil {
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	return string(body), nil
}

func (client *PlaylyfeClient) Delete(endpoint Endpoint, dataStructResult interface{}) (err error) {
	body, err := client.DeleteRaw(endpoint)

	err = json.Unmarshal([]byte(body), &dataStructResult)

	return
}
