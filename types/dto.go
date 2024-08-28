package vgtypes

import "github.com/goccy/go-json"

type DTOResponse interface {
	IsDTOResponse()
}

type DTORequest interface {
	IsDTORequest()
}

type DataRequest struct{}

func (d *DataRequest) IsDTORequest() {}

type DataResponse struct{}

func (d *DataResponse) IsDTOResponse() {}

type Update struct{}

type ReponseWrapper struct {
	DataResponse
	Response json.RawMessage `json:"response"`
	Error    json.RawMessage `json:"error"`
}

type GroupsGetLongPollServerRequest struct {
	DataRequest
	GroupID int64 `json:"group_id"`
}

type GroupsGetLongPollServerResponse struct {
	DataResponse
	Key       string `json:"key"`
	Server    string `json:"server"`
	Timestamp int    `json:"ts"`
}
