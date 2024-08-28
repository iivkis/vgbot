package vgstd

import (
	"context"

	vgtypes "github.com/iivkis/vgbot/types"
)

func MessageFilter() vgtypes.UpdateFilter {
	return func(ctx context.Context, update vgtypes.Update) bool {
		return true
	}
}
