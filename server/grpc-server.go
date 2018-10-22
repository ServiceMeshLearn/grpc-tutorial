package main

import (
	"fmt"
	"log"
	"net"

	"github.com/liuliqiang/grpc-example/proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var host = "localhost"
var port = 9960

func main() {
	listen, err := net.Listen("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := &server{}
	grpcServ := grpc.NewServer()
	api.RegisterPingServer(grpcServ, s)

	if err := grpcServ.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type server struct {
}

func (*server) SayHello(ctx context.Context, req *api.PingMessage) (*api.PingMessage, error) {
	log.Printf("[I] recv hello message: %s", req.Message)
	req.Message = fmt.Sprintf("Yes, i have got: %s", req.Message)
	return req, nil
}

func NewServer() api.PingServer {
	return &server{}
}
