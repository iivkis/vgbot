package vgtypes

import "context"

type Update struct{}

type UpdateFilter func(ctx context.Context, update Update) bool

type UpdateHandlerFunc func(ctx context.Context, update Update)
