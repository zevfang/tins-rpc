package common

import (
	"context"
	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"time"
)

const (
	RPCX = "RPCx"
	GRPC = "gRPC"
)

type RequestData struct {
	Address       string
	ServicePath   string
	ServiceMethod string
	Metadata      map[string]string
	Payload       []byte
}

func Call(entity string, req RequestData) (header map[string]string, body []byte, err error) {
	switch entity {
	case RPCX:
		return DoRPCX(req)
	case GRPC:
		return DoGRPC(req)
	}
	return
}

func DoRPCX(req RequestData) (header map[string]string, body []byte, err error) {
	ctx := context.Background()
	option := client.DefaultOption
	option.Heartbeat = true
	option.HeartbeatInterval = 1 * time.Second
	option.TCPKeepAlivePeriod = 1 * time.Second
	rClient := client.NewClient(option)
	err = rClient.Connect(`tcp`, req.Address)
	if err != nil {
		return nil, nil, err
	}
	request := protocol.NewMessage()
	request.SetMessageType(protocol.Request)
	request.SetSerializeType(protocol.JSON)
	request.ServicePath = req.ServicePath
	request.ServiceMethod = req.ServiceMethod
	request.Metadata = req.Metadata
	request.Payload = req.Payload
	header, body, err = rClient.SendRaw(ctx, request)
	return
}

func DoGRPC(req RequestData) (header map[string]string, body []byte, err error) {
	return
}
