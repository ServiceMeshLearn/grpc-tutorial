package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/liuliqiang/grpc-example/proto"
	"github.com/soheilhy/cmux"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var host = "localhost"
var port = 9960

func main() {
	addr := fmt.Sprintf("%s:%d", host, port)
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener) })
	g.Go(func() error { return m.Serve() })

	log.Println("run server:", g.Wait())
}

func httpServe(l net.Listener) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	grpcEp := fmt.Sprintf("%s:%d", host, port)
	err := api.RegisterPingHandlerFromEndpoint(ctx, mux, grpcEp, opts)
	if err != nil {
		log.Fatalf("Failed to registe: %v", err)
	}

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}

func grpcServe(l net.Listener) error {
	s := &server{}
	grpcServ := grpc.NewServer()
	api.RegisterPingServer(grpcServ, s)

	return grpcServ.Serve(l)
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
