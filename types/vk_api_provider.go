package vgtypes

type VKAPIProvider interface {
	Call(method string, data DataEncoder, dst DataDecoder) error
}
