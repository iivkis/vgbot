package vgbot

import (
	vgtypes "github.com/iivkis/vgbot/types"
)

type VKMessagesAPI struct {
}

func NewVKMessagesAPI() vgtypes.VKMessagesAPI {
	return &VKMessagesAPI{}
}

func (v *VKMessagesAPI) Delete(peerID int64, messageID int) error {
	panic("unimplemented")
}

func (v *VKMessagesAPI) Send(peerID int64, message string) error {
	panic("unimplemented")
}

type VKAPI struct {
	token string
}

func NewVKAPI(token string) vgtypes.VKAPI {
	return &VKAPI{
		token: token,
	}
}

func (a *VKAPI) Messages() vgtypes.VKMessagesAPI {
	return NewVKMessagesAPI()
}
