package vgtypes

type VKGroupsAPI interface {
	GetLongPollServer(groupID int64) (*GroupsGetLongPollServerResponse, error)
}

type VKMessagesAPI interface {
	Send(peerID int64, message string) error
	Delete(peerID int64, messageID int) error
}

type VKAPI interface {
	Groups() VKGroupsAPI
	Messages() VKMessagesAPI
}
