package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	greeterPb "tins-rpc/example/proto"
)

const (
	port = ":9080"
)

type GreeterService struct {
	greeterPb.UnimplementedGreeterServer
}

func (s *GreeterService) SayHello(ctx context.Context, in *greeterPb.HelloRequest) (*greeterPb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	msg := in.GetName() + " say hello for gRPC"
	reply := &greeterPb.HelloReply{
		Message: msg,
	}
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	srv := grpc.NewServer()
	greeterPb.RegisterGreeterServer(srv, &GreeterService{})
	log.Println("gRPC server is running...")
	if err := srv.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
