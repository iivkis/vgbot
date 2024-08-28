package main

import (
	"context"

	vgstd "github.com/iivkis/vgbot/std"
	vgtypes "github.com/iivkis/vgbot/types"
)

type BasicHandler struct{}

func NewBasicHandler() vgtypes.RouteAdjuster {
	return &BasicHandler{}
}

func (b *BasicHandler) Setup(r vgtypes.Router) {
	r.On(
		vgstd.MessageFilter(),
	)(func(ctx context.Context, update vgtypes.Update) {
		println("hello world! :D")
	})
}
