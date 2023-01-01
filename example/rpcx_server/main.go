package main

import (
	"context"
	"fmt"
	"github.com/smallnest/rpcx/server"
	"log"
	greeterPb "tins-rpc/example/proto"
)

const (
	port = ":9081"
)

type GreeterService struct{}

func (s *GreeterService) SayHello(ctx context.Context, in *greeterPb.HelloRequest, reply *greeterPb.HelloReply) error {
	log.Printf("Received: %v", in.GetName())
	msg := in.GetName() + " say hello for RPCx"
	reply.Message = msg
	return nil
}

func Router(router *server.Server) {
	err := router.RegisterName("Greeter", new(GreeterService), "")
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// TLS
	//cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	//if err != nil {
	//log.Print(err)
	//return
	//}
	//config := &tls.Config{Certificates: []tls.Certificate{cert}}
	//s := server.NewServer(server.WithTLSConfig(config))

	s := server.NewServer()
	s.DisableHTTPGateway = true
	s.DisableJSONRPC = true
	Router(s)
	log.Println("RPCx server is running...")
	log.Fatal(s.Serve(`tcp`, fmt.Sprintf("%s", port)))
}
