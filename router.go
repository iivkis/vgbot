package vgbot

import (
	vgtypes "github.com/iivkis/vgbot/types"
)

type RouteHandler struct {
	filters     []vgtypes.UpdateFilter
	handlerFunc vgtypes.UpdateHandlerFunc
}

type Router struct {
	middlewares []vgtypes.Middleware
	routers     []vgtypes.Router
	handlers    []*RouteHandler
}

func NewRouter() vgtypes.Router {
	return &Router{}
}

func (r *Router) Handle(update vgtypes.Update) {
	panic("unimplemented")
}

func (r *Router) IncludeRouter(router vgtypes.Router) {
	panic("unimplemented")
}

func (r *Router) IncludeAdjuster(adjuster ...vgtypes.RouteAdjuster) {
	panic("unimplemented")
}

func (r *Router) On(filters ...vgtypes.UpdateFilter) func(vgtypes.UpdateHandlerFunc) {
	return func(f vgtypes.UpdateHandlerFunc) {
		r.handlers = append(r.handlers, &RouteHandler{
			filters:     filters,
			handlerFunc: f,
		})
	}
}
