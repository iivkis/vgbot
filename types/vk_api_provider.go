package vgtypes

type VKAPIProvider interface {
	Call(method string, data DataEncoder) (DataDecoder, error)
}
