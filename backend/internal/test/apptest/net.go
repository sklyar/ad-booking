package apptest

import (
	"connectrpc.com/connect"
)

type newServiceClientFunc[T any] func(httpClient connect.HTTPClient, baseURL string, opts ...connect.ClientOption) T

func NewGRPCClient[ServiceClient any](suite *Suite, fn newServiceClientFunc[ServiceClient]) ServiceClient {
	return fn(suite.client, suite.serverBaseURL)
}
