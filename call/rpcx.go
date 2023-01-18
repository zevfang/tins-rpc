package call

import (
	"context"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
)

func DoRPCX(req RequestData) (header map[string]string, body []byte, err error) {
	ctx := context.Background()
	option := client.DefaultOption
	// tls
	//conf := &tls.Config{
	//	InsecureSkipVerify: true,
	//}
	//option.TLSConfig = conf
	option.Heartbeat = true
	option.HeartbeatInterval = 9999 * time.Second
	option.TCPKeepAlivePeriod = 9999 * time.Second
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
