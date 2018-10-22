package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	"github.com/liuliqiang/grpc-example/proto"
)

var host = "localhost"
var port = 9960

func main() {
	addr := fmt.Sprintf("%s:%d", host, port)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	cli := api.NewPingClient(conn)
	resp, err := cli.SayHello(context.Background(), &api.PingMessage{Message: "Hello"})
	if err != nil {
		log.Fatalf("Failed to say hello: %v", err)
	}
	log.Printf("[I] Server response: %s", resp.Message)
}
