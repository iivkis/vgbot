package vgbot

import (
	"fmt"

	vgtypes "github.com/iivkis/vgbot/types"
)

type VKBotAPI struct {
	token string
}

func NewVKBotAPI(token string) vgtypes.VKBotAPI {
	return &VKBotAPI{
		token: token,
	}
}

// PrintHello implements vgtypes.VKBotAPI.
func (v *VKBotAPI) PrintHello() {
	fmt.Println("Hello from VKBot API! :D")
}
