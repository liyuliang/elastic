package elastic

import (
	"context"
	"errors"
	"net/http"
	"strings"
)

type bulk struct {
	_client  *client
	index    string
	typeName string
	requests []bulkRequest
}

func (b *bulk) Add(requests ...bulkRequest) *bulk {

	for _, request := range requests {
		b.requests = append(b.requests, request)
	}
	return b
}

func (b *bulk) Index(index string) *bulk {
	b.index = index
	return b
}

func (b *bulk) Type(typeName string) *bulk {
	b.typeName = typeName
	return b
}

func (b *bulk) Do(ctx context.Context) *esResponse {
	resp := new(esResponse)

	if len(b.requests) == 0 {
		resp.Err = errors.New("elastic: No bulk actions to commit")
		return resp
	}

	body, err := b.bodyAdString()
	if err != nil {
		resp.Err = err
		return resp
	}

	url := b._client.host + "/" + b.index + "/" + b.typeName+"/_bulk"
	httpClient := &http.Client{}
	request, _ := http.NewRequest("POST", url, strings.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	httpResp, err := httpClient.Do(request)

	if err != nil {
		resp.Err = err
		return resp
	}
	resp = response(httpResp, err)
	return resp
}

func (b *bulk) bodyAdString() (body string, err error) {

	for _, req := range b.requests {

		sources, err := req.Source()
		if err != nil {
			return "", err
		}

		for _, source := range sources {
			body += source + "\n"
		}
	}

	return body, nil
}
