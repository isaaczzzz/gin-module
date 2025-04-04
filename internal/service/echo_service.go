package service

import (
	"context"
	"github.com/isaaczzzz/gin-module/pkg/rpcclient"
	"github.com/isaaczzzz/gin-module/rpc_gen/api/echo"
	"log"
	"time"
)

type IEchoService interface {
	Echo(ctx context.Context, message string) (string, error)
}

type EchoService struct {
	echoClient echo.EchoServiceClient
}

func newEchoService(client *rpcclient.EchoClient) IEchoService {
	return &EchoService{
		echoClient: echo.NewEchoServiceClient(client.Conn),
	}
}

func (s *EchoService) Echo(ctx context.Context, message string) (string, error) {
	grpcCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	req := &echo.EchoRequest{
		Message: message,
	}

	resp, err := s.echoClient.Echo(grpcCtx, req)
	if err != nil {
		log.Printf("Failed to call gRPC Echo service: %v", err)
		return "", err
	}
	return resp.Message, nil
}
