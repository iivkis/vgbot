package vgtypes

type DataEncoder interface {
	Encode() ([]byte, error)
}

type DataDecoder interface {
	Decode(data []byte) error
}

type VKAPIProvider interface {
	Call(method string, data DataEncoder) (DataDecoder, error)
}

type VKGroupsAPI interface {
	GetLongPollServer(groupID int) GroupsGetLongPollServerResponse
}

type VKMessagesAPI interface {
	Send(peerID, message string) error
	Delete(peerID, messageID int) error
}

type VKAPI interface {
	Messages() VKMessagesAPI
}
