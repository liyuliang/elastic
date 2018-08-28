package elastic

import (
	"strings"
	"encoding/json"
)

func request(data map[string]string) *strings.Reader {
	body := ""
	if len(data) > 0 {
		byteData, _ := json.Marshal(data)
		body = string(byteData)
	}
	return strings.NewReader(body)

}
