package main

import (
	"context"

	vgstd "github.com/iivkis/vgbot/std"
	vgtypes "github.com/iivkis/vgbot/types"
)

type BasicHandler struct {
	vkapi vgtypes.VKAPI
}

func NewBasicHandler(
	vkapi vgtypes.VKAPI,
) vgtypes.RouteHandler {
	return &BasicHandler{
		vkapi: vkapi,
	}
}

func (h *BasicHandler) Setup(r vgtypes.Router) {
	r.On(
		vgstd.MessageFilter(),
	)(func(ctx context.Context, update vgtypes.Update) {
		h.vkapi.Messages().Send("", "")
	})
}
