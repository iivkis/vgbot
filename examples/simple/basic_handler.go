package main

import (
	"context"

	vgstd "github.com/iivkis/vgbot/std"
	vgtypes "github.com/iivkis/vgbot/types"
)

type BasicHandler struct {
	botapi vgtypes.VKBotAPI
}

func NewBasicHandler(
	botapi vgtypes.VKBotAPI,
) vgtypes.RouteAdjuster {
	return &BasicHandler{
		botapi: botapi,
	}
}

func (h *BasicHandler) Setup(r vgtypes.Router) {
	r.On(
		vgstd.MessageFilter(),
	)(func(ctx context.Context, update vgtypes.Update) {
		h.botapi.PrintHello()
	})
}
