package call

import (
	"context"
	"encoding/json"
	"fmt"

	"tins-rpc/common"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

var unaryStreamDesc = &grpc.StreamDesc{ServerStreams: false, ClientStreams: false}

func DoGRPC(req RequestData) (header map[string]string, body []byte, err error) {
	//set metadata
	md := metadata.New(req.Metadata)
	ctx := metadata.NewOutgoingContext(context.Background(), md)
	//dial
	conn, err := grpc.Dial(req.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}
	defer conn.Close()
	// method = "/greeter.Greeter/SayHello"
	cs, err := conn.NewStream(ctx, unaryStreamDesc,
		fmt.Sprintf("/%s.%s/%s", req.PackageName, req.ServicePath, req.ServiceMethod))
	if err != nil {
		return nil, nil, err
	}
	// messageName = "greeter.HelloRequest"
	reqData, err := common.JsonToPb(req.Fd.FileDescriptor, req.Fd.Request, req.Payload)
	if err != nil {
		return nil, nil, err
	}
	if err := cs.SendMsg(reqData); err != nil {
		return nil, nil, err
	}
	// messageName = "greeter.HelloReply"
	out, err := common.JsonToPbDynamic(req.Fd.FileDescriptor, req.Fd.Return)
	if err != nil {
		return nil, nil, err
	}
	err = cs.RecvMsg(out)
	if err != nil {
		return nil, nil, err
	}
	resp, err := json.Marshal(out)
	if err != nil {
		return nil, nil, err
	}
	fmt.Println(string(resp))
	return nil, resp, nil
}
