package elastic

import "encoding/json"

func BulkUpdateRequest() *bulkUpdateRequest {
	req := new(bulkUpdateRequest)
	return req
}


type bulkUpdateRequest struct {
	bulkRequest
	index string
	typ string
	id string
	doc interface{}
}

type bulkUpdateRequestCommandOp struct {
	Index string `json:"_index,omitempty"`
	Type string `json:"_type,omitempty"`
	Id string `json:"_id,omitempty"`
}


func (req *bulkUpdateRequest) Id(id string) *bulkUpdateRequest {
	req.id = id
	return req

}
func (req *bulkUpdateRequest) Doc(doc interface{}) *bulkUpdateRequest {
	req.doc = doc
	return req
}

func (req *bulkUpdateRequest) OpType() string {
	return "update"
}

func (req *bulkUpdateRequest) Source() ([]string, error) {

	lines := make([]string, 2)

	indexCommand := bulkUpdateRequestCommandOp{
		Index: req.index,
		Type:  req.typ,
		Id:    req.id,
	}

	command := make(map[string]bulkUpdateRequestCommandOp)
	command[req.OpType()] = indexCommand

	body, err := json.Marshal(command)

	if err != nil {
		return nil, err
	}

	lines[0] = string(body)
	lines[1], err = req.docAsString()

	if err != nil {
		return nil, err
	}

	return lines, nil
}

func (req *bulkUpdateRequest) docAsString() (string, error) {
	body, err := json.Marshal(req.doc)
	return string(body), err
}
