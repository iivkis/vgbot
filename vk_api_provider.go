package vgbot

import (
	"net/http"

	vgtypes "github.com/iivkis/vgbot/types"
)

type VKAPIProvider struct {
	token   string
	client  http.Client
	baseURL string
}

func NewVKAPIProvider(token string) vgtypes.VKAPIProvider {
	return &VKAPIProvider{
		token:   token,
		client:  http.Client{},
		baseURL: "https://api.vk.com/method/",
	}
}

// Call implements vgtypes.VKAPIProvider.
func (v *VKAPIProvider) Call(method string, data vgtypes.DataEncoder) ([]byte, error) {
	panic("unimplemented")
}
