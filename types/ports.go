package vgtypes

import "context"

type VKBotAPI interface {
}

type Middleware interface {
	Call(ctx context.Context, update Update) bool
}

type RouteAdjuster interface {
	Setup(r Router)
}

type Router interface {
	On(filters ...UpdateFilter) func(UpdateHandlerFunc)
	IncludeRouter(router Router)
	IncludeAdjuster(adjuster ...RouteAdjuster)
	Handle(update Update)
}

type UpdateListener interface {
	Listen()
}
