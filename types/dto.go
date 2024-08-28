package vgtypes

type Update struct{}

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
