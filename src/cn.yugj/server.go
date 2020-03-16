package main

import (
	"github.com/openzipkin/zipkin-go"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	logreporter "github.com/openzipkin/zipkin-go/reporter/log"
)

const (
	PORT = ":50001"
)

type server struct {}

func (s *server) SayHello(ctx context.Context, in *HelloRequest) (*HelloReply, error) {
	log.Println("request: ", in.Name)
	return &HelloReply{Message: "Hello " + in.Name}, nil
}

func callJava()  {

}

func main() {

	lis, err := net.Listen("tcp", PORT)

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	//create server with zipkin wrapper
	reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	tracer, err := zipkin.NewTracer(reporter)
	server2 := grpc.NewServer(grpc.StatsHandler(zipkingrpc.NewServerHandler(tracer)))

	//s := grpc.NewServer()

	RegisterGreeterServer(server2, &server{})

	log.Println("rpc服务已经开启")
	server2.Serve(lis)
}