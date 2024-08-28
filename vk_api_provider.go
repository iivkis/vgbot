package vgbot

import (
	"fmt"
	"io"
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
func (v *VKAPIProvider) Call(method string, data vgtypes.DTORequest, dst vgtypes.DTOResponse) error {
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

	var resWrap vgtypes.ReponseWrapper
	if err := json.Unmarshal(b, &resWrap); err != nil {
		return fmt.Errorf("failed to unmarshal response: %w", err)
	}

	if resWrap.Error != nil {
		return fmt.Errorf("VKAPIError: %s", string(resWrap.Error))
	}

	if err := json.Unmarshal(resWrap.Response, dst); err != nil {
		return fmt.Errorf("failed to unmarshal response data: %w", err)
	}

	return nil
}
