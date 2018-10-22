package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/liuliqiang/grpc-example/proto"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var host = "localhost"
var grpcPort = 9960
var port = 9961

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcEp := fmt.Sprintf("%s:%d", host, grpcPort)
	err := api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcEp, opts)
	if err != nil {
		log.Fatalf("Failed to registe: %v", err)
	}

	listenAddr := fmt.Sprintf("%s:%d", host, port)
	if err = http.ListenAndServe(listenAddr, mux); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
