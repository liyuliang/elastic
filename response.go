package elastic

import (
	"net/http"
	"strings"
	"io/ioutil"
	"errors"
)

type esResponse struct {
	Content    string
	Header     map[string]string
	StatusCode int
	Err        error
}

func (r *esResponse) Headers() (headers string) {

	if len(r.Header) != 0 {

		for key, value := range r.Header {
			headers += key + ":" + value + ","
		}
	}
	return headers
}

func response(httpResp *http.Response, err error) *esResponse {

	resp := new(esResponse)
	if httpResp == nil {
		resp.Err = errors.New("ES Response Error")
		return resp
	}
	resp.Err = err
	resp.Header = make(map[string]string)
	for key, value := range httpResp.Header {
		resp.Header[key] = strings.Join(value[:], ";")
	}

	if err != nil {
		resp.Err = err

	} else {
		defer httpResp.Body.Close()

		contents, err := ioutil.ReadAll(httpResp.Body)

		resp.StatusCode = httpResp.StatusCode
		resp.Err = err
		resp.Content = string(contents)
	}
	return resp
}
