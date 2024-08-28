package vgtypes

import "github.com/goccy/go-json"

type Update struct{}

type ReponseWrapper struct {
	DataObject
	Response json.RawMessage `json:"response"`
	Error    json.RawMessage `json:"error"`
}

func (r *ReponseWrapper) Decode(data []byte) error {
	return json.Unmarshal(data, r)
}

type GroupsGetLongPollServerRequest struct {
	DataObject
	GroupID int64 `json:"group_id"`
}

type GroupsGetLongPollServerResponse struct {
	DataObject
	Key       string `json:"key"`
	Server    string `json:"server"`
	Timestamp int    `json:"ts"`
}
