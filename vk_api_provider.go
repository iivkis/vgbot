package vgbot

import (
	"net/http"

	"github.com/goccy/go-json"
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
func (v *VKAPIProvider) Call(method string, data vgtypes.DataEncoder, dst vgtypes.DataDecoder) error {
	req, err := http.NewRequest(http.MethodPost, v.baseURL+method, nil)
	if err != nil {
		return err
	}

	res, err := v.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	var response vgtypes.ReponseWrapper
	if err := json.NewDecoder(req.Body).Decode(&response); err != nil {
		return err
	}

	if err := dst.Decode(response.Response); err != nil {
		return err
	}

	return nil
}
