package togotime

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var baseURL string = "https://www.toggl.com/api/v8/"

// HTTPHelper A class to wrap http requests and functions for interaction with the toggl api
type HTTPHelper struct {
	APIToken string
}

// GetRawJSON A Wrapper for net/http requests to return byte array of json response
func (h HTTPHelper) GetRawJSON(url string) []byte {
	req := h.createRequest("GET", url)
	req.SetBasicAuth(h.APIToken, "api_token")
	resp := h.makeRequest(req)
	raw := h.readBodyToBytes(resp.Body)
	// TODO: Actually needs to do something about 400 codes
	return raw
}

// makeRequest Makes a http request with the DefaultClient using a given made request and logs errors
func (h HTTPHelper) makeRequest(req *http.Request) *http.Response {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	return resp
}

// createRequest Creates a request and logs errors
func (h HTTPHelper) createRequest(method string, url string) *http.Request {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}
	return req
}

// readBodyToBytes reads a response from
func (h HTTPHelper) readBodyToBytes(respBody io.ReadCloser) []byte {
	body, err := ioutil.ReadAll(respBody)
	if err != nil {
		log.Fatal(err)
	}
	return body
}
