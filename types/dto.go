package vgtypes

import "github.com/goccy/go-json"

type Update struct{}

type ReponseWrapper struct {
	Response json.RawMessage `json:"response"`
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
