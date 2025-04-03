package rpcclient

import (
	"context"
	"github.com/isaaczzzz/gin-module/rpc_gen/api/echo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// EchoClient 封装 Echo 服务的 gRPC 客户端
type EchoClient struct {
	Client echo.EchoServiceClient
	Conn   *grpc.ClientConn
}

// NewEchoClient 创建一个新的 Echo 服务 gRPC 客户端
func NewEchoClient(address string) (*EchoClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	client := echo.NewEchoServiceClient(conn)
	return &EchoClient{
		Client: client,
		Conn:   conn,
	}, nil
}

// Echo 调用 Echo 服务的 Echo 方法
func (c *EchoClient) Echo(ctx context.Context, message string) (*echo.EchoResponse, error) {
	req := &echo.EchoRequest{
		Message: message,
	}
	return c.Client.Echo(ctx, req)
}

// Close 关闭 gRPC 连接
func (c *EchoClient) Close() error {
	return c.Conn.Close()
}
