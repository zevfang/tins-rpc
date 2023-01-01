package example

//go:generate protoc -I=. --go_out=. --go-grpc_out=. greeter.proto
