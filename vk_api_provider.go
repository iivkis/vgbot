package vgbot

import (
	"fmt"
	"io"
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

	b, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	resWrap := &vgtypes.ReponseWrapper{}
	if err := resWrap.Decode(b); err != nil {
		return err
	}

	fmt.Println("434233432242342342423423")

	if err := dst.Decode(resWrap.Response); err != nil {
		return err
	}

	return nil
}
