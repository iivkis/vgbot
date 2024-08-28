package vgtypes

import "github.com/goccy/go-json"

type DataEncoder interface {
	Encode() ([]byte, error)
}

type DataDecoder interface {
	Decode(data []byte) error
}

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
