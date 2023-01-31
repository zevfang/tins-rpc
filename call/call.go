package call

import "github.com/jhump/protoreflect/desc"

type FrameTypeEnum string

const (
	RPCX FrameTypeEnum = "RPCx"
	GRPC FrameTypeEnum = "gRPC"
)

var FrameTypes = []string{RPCX.ToString(), GRPC.ToString()}

func (b FrameTypeEnum) ToString() string {
	return string(b)
}

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
	case RPCX.ToString():
		return DoRPCX(req)
	case GRPC.ToString():
		return DoGRPC(req)
	}
	return
}
