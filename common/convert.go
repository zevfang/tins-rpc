package common

import (
	"github.com/golang/protobuf/proto"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/dynamic"
)

func GetProtoFileDescriptor(path string) *desc.FileDescriptor {
	p := protoparse.Parser{}
	fds, err := p.ParseFiles(path)
	if err != nil {
		return nil
	}
	fd := fds[0]
	return fd
}

func JsonToPb(fd *desc.FileDescriptor, messageName string, jsonStr []byte) (proto.Message, error) {
	msg := fd.FindMessage(messageName)
	dymsg := dynamic.NewMessage(msg)
	err := dymsg.UnmarshalJSON(jsonStr)
	if err != nil {
		return nil, nil
	}
	return dymsg, nil
}

func JsonToPbDynamic(fd *desc.FileDescriptor, messageName string) (proto.Message, error) {
	msg := fd.FindMessage(messageName)
	dymsg := dynamic.NewMessage(msg)
	return dymsg, nil
}

func PbToJson(fd *desc.FileDescriptor, messageName string, protoData []byte) ([]byte, error) {
	msg := fd.FindMessage(messageName)
	dymsg := dynamic.NewMessage(msg)
	err := proto.Unmarshal(protoData, dymsg)
	jsonByte, err := dymsg.MarshalJSON()
	return jsonByte, err
}
