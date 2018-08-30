package elastic

import "encoding/json"

func BulkIndexRequest() *bulkIndexRequest {
	req := new(bulkIndexRequest)
	return req
}

type bulkIndexRequest struct {
	bulkRequest
	index string
	typ   string
	id    string
	doc   interface{}
}

type bulkIndexRequestCommandOp struct {
	Index string `json:"_index,omitempty"`
	Type  string `json:"_type,omitempty"`
	Id    string `json:"_id,omitempty"`
}

func (req *bulkIndexRequest) Id(id string) *bulkIndexRequest {
	req.id = id
	return req
}

func (req *bulkIndexRequest) Doc(doc interface{}) *bulkIndexRequest {
	req.doc = doc
	return req
}

func (req *bulkIndexRequest) OpType() string {
	return "index"
}

func (req *bulkIndexRequest) Source() ([]string, error) {

	lines := make([]string, 2)

	indexCommand := bulkIndexRequestCommandOp{
		//Index: req.index,
		//Type:  req.typ,
		Id:    req.id,
	}

	command := make(map[string]bulkIndexRequestCommandOp)
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

func (req *bulkIndexRequest) docAsString() (string, error) {
	body, err := json.Marshal(req.doc)
	return string(body), err
}
