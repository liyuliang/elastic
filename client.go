package elastic

import (
	"net/http"
	"io"
)

func Client(host string) *client {
	c := new(client)
	c.host = host
	return c
}

type client struct {
	host string
	bulk *bulk
}

func (c *client) Bulk() *bulk {

	if c.bulk == nil {
		c.bulk = new(bulk)
		c.bulk._client = c
	}
	return c.bulk
}

func (c *client) Url(uri string) string {
	return c.host + uri
}

func (c *client) Put(uri string, body io.Reader) *esResponse {

	httpClient := &http.Client{}
	request, _ := http.NewRequest("PUT", c.Url(uri), body)
	request.Header.Add("Accept", "application/json")
	request.Header.Set("Content-Type", "application/json")
	httpResp, err := httpClient.Do(request)

	resp := response(httpResp, err)
	return resp
}

/**
 * Not specify the id
 */
func (c *client) Post(uri string, body io.Reader) *esResponse {

	httpClient := &http.Client{}
	request, _ := http.NewRequest("POST", c.Url(uri), body)
	request.Header.Set("Content-Type", "application/json")
	httpResp, err := httpClient.Do(request)

	resp := response(httpResp, err)
	return resp
}

func (c *client) Get(uri string) *esResponse {

	httpClient := &http.Client{}
	request, _ := http.NewRequest("GET", c.Url(uri), nil)
	httpResp, err := httpClient.Do(request)
	resp := response(httpResp, err)
	return resp
}

func (c *client) Delete(uri string) *esResponse {
	httpClient := &http.Client{}
	request, _ := http.NewRequest("DELETE", c.Url(uri), nil)
	httpResp, err := httpClient.Do(request)
	resp := response(httpResp, err)
	return resp
}
