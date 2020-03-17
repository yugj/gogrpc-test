package main

import (
	"github.com/openzipkin/zipkin-go"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	zipkingrpc "github.com/openzipkin/zipkin-go/middleware/grpc"
	logreporter "github.com/openzipkin/zipkin-go/reporter/log"
)

const (
	address = "localhost:50001"
)

func main() {

	reporter := logreporter.NewReporter(log.New(os.Stderr, "", log.LstdFlags))
	tracer, _ := zipkin.NewTracer(reporter)

	//conn, err := grpc.Dial(address, grpc.WithInsecure())

	var conn, err = grpc.Dial(address, grpc.WithInsecure(),grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)))

	//conn, err = grpc.Dial(address, grpc.WithInsecure(), grpc.WithStatsHandler(zipkingrpc.NewClientHandler(tracer)))

	//with remote service todo to test
	//var conn, err = grpc.Dial(
	//	address,
	//	grpc.WithInsecure(),
	//	grpc.WithStatsHandler(zipkingrpc.NewClientHandler(
	//		tracer,
	//		zipkingrpc.WithRemoteServiceName("remoteService"))))


	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	c := NewGreeterClient(conn)

	name := "hello xx"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SayHello(context.Background(), &HelloRequest{Name: name})

	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(r.Message)
}