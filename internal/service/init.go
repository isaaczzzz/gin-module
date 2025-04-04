package service

import "github.com/isaaczzzz/gin-module/pkg/rpcclient"

type Services struct {
	EchoService IEchoService
}

func InitServices(clients *rpcclient.RPCClients) *Services {
	return &Services{
		EchoService: newEchoService(clients.EchoClient),
	}
}
