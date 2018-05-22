package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	pb "self/test/example/proto"
)

const (
	Address = "127.0.0.1:50052"
)

// 实现了grpc定义的方法
type helloService struct{}

var HelloService helloService

func (h helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (resp *pb.HelloReply, err error) {
	resp = &pb.HelloReply{}
	resp.Message = fmt.Sprint("welcome to snaplingo ", in.Name, " !!!")
	return
}

func main() {
	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("failed to connect %v", err)
	}

	// 实例化一个grpc Server
	/*
		实例话一个server,但是它还没有服务注册，并且还没有开启接收请求.
	*/
	s := grpc.NewServer()

	// 注册HelloService--在调用服务之前必须被调用的(是由ide产生的)--注册一个服务
	pb.RegisterHelloServer(s, HelloService)
	fmt.Println("listen on ", Address)
	s.Serve(listen)

}
