package vgtypes

import "context"

type UpdateFilter func(ctx context.Context, update Update) bool

type UpdateHandlerFunc func(ctx context.Context, update Update)

type UpdateListener interface {
	Listen()
}
