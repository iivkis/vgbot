package vgbot

import vgtypes "github.com/iivkis/vgbot/types"

type VKBotAPI struct {
	token string
}

func NewVKBotAPI(token string) vgtypes.VKBotAPI {
	return &VKBotAPI{
		token: token,
	}
}
