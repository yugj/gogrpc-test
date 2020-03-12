package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
)

const (
	PORT = ":50001"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Println("request: ", in.Name)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	RegisterGreeterServer(s, &server{})

	log.Println("rpc服务已经开启")
	s.Serve(lis)
}