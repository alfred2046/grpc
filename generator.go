// go:generate protoc --go_out=. --go-grpc_out=. ./sample.proto
package grpc

func Gen() string {
	return "message from generator"
}
