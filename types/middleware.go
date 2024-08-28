package vgtypes

import "context"

type Middleware interface {
	Call(ctx context.Context, update Update) bool
}
