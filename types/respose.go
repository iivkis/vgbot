package vgtypes

import "github.com/goccy/go-json"

var (
	_ DataEncoder = &DataObject{}
	_ DataDecoder = &DataObject{}
)

type DataObject struct{}

func (d *DataObject) Encode() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DataObject) Decode(data []byte) error {
	return json.Unmarshal(data, d)
}

type Update struct{}

type GroupsGetLongPollServerResponse struct {
	DataObject
	Key       string `json:"key"`
	Server    string `json:"server"`
	Timestamp int    `json:"ts"`
}
