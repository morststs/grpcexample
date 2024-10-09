package main

import (
	"example/grpc_sample"
	"log"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func main() {
	log.Print("main start")

	if len(os.Args) != 2 {
		log.Fatalf("One argument is required")
	}

	listen, err := net.Listen("tcp", os.Args[1])
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	grpc_sample.RegisterSampleServiceServer(grpcServer, &Sample{})

	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

	log.Print("main end")
}

type Sample struct {
	name string
}

func (s *Sample) GetData(
	ctx context.Context,
	message *grpc_sample.Message,
) (*grpc_sample.Message, error) {
	log.Print(message)
	return message, nil
}
