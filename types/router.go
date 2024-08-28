package vgtypes

type RouteHandler interface {
	Setup(r Router)
}

type Router interface {
	On(filters ...UpdateFilter) func(UpdateHandlerFunc)
	IncludeRouter(router Router)
	IncludeHandler(handlers ...RouteHandler)
	Handle(update Update)
}
