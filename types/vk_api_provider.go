package vgtypes

type VKAPIProvider interface {
	Call(method string, data DTORequest, dst DTOResponse) error
}
