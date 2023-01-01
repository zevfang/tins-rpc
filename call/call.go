package call

import "github.com/jhump/protoreflect/desc"

const (
	RPCX = "RPCx"
	GRPC = "gRPC"
)

type RequestData struct {
	Fd            ProtoDescriptor
	Address       string
	PackageName   string
	ServicePath   string
	ServiceMethod string
	Metadata      map[string]string
	Payload       []byte
}

type ProtoDescriptor struct {
	FileDescriptor *desc.FileDescriptor
	Request        string
	Return         string
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
