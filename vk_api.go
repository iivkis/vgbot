package vgbot

import (
	vgtypes "github.com/iivkis/vgbot/types"
)

// VKGroupsAPI
type VKGroupsAPI struct {
	provider vgtypes.VKAPIProvider
}

func NewVKGroupsAPI(provider vgtypes.VKAPIProvider) vgtypes.VKGroupsAPI {
	return &VKGroupsAPI{
		provider: provider,
	}
}

func (v *VKGroupsAPI) GetLongPollServer(groupID int64) (*vgtypes.GroupsGetLongPollServerResponse, error) {
	res, err := v.provider.Call(
		"groups.getLongPollServer",
		&vgtypes.GroupsGetLongPollServerRequest{
			GroupID: groupID,
		})
	if err != nil {
		return nil, err
	}

	data := &vgtypes.GroupsGetLongPollServerResponse{}
	if err := data.Decode(res); err != nil {
		return nil, err
	}

	return data, nil
}

// MessagesAPI
type VKMessagesAPI struct {
	provider vgtypes.VKAPIProvider
}

func NewVKMessagesAPI(provider vgtypes.VKAPIProvider) vgtypes.VKMessagesAPI {
	return &VKMessagesAPI{
		provider: provider,
	}
}

func (v *VKMessagesAPI) Delete(peerID int64, messageID int) error {
	panic("unimplemented")
}

func (v *VKMessagesAPI) Send(peerID int64, message string) error {
	panic("unimplemented")
}

// VKAPI
type VKAPI struct {
	provider    vgtypes.VKAPIProvider
	messagesAPI vgtypes.VKMessagesAPI
	groupsAPI   vgtypes.VKGroupsAPI
}

func NewVKAPI(provider vgtypes.VKAPIProvider) vgtypes.VKAPI {
	return &VKAPI{
		provider:    provider,
		messagesAPI: NewVKMessagesAPI(provider),
	}
}

func (a *VKAPI) Groups() vgtypes.VKGroupsAPI {
	return a.groupsAPI
}

func (a *VKAPI) Messages() vgtypes.VKMessagesAPI {
	return a.messagesAPI
}
