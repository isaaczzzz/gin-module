package rpcclient

import (
	"fmt"
	"github.com/isaaczzzz/gin-module/pkg/config"
	"log"
)

type RPCClients struct {
	EchoClient *EchoClient
}

func InitRPCClients(cfg *config.Config) *RPCClients {
	echoClient, err := NewEchoClient(fmt.Sprintf("%s:%s", cfg.GRPC.Address, cfg.GRPC.Port))
	if err != nil {
		log.Fatalf("Failed to create EchoClient: %v", err)
	}
	return &RPCClients{
		EchoClient: echoClient,
	}
}

func (r *RPCClients) CloseRPCClients() error {
	err := r.EchoClient.Close()
	if err != nil {
		return err
	}
	return nil
}
