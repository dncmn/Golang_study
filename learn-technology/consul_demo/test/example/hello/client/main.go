package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "self/test/example/proto"
)

const (
	Address = "127.0.0.1:50052"
)

func main() {
	conn, err := grpc.Dial(Address, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	repBody := new(pb.HelloRequest)
	repBody.Name = "SanpLingo"

	r, err := c.SayHello(context.Background(), repBody)
	if err != nil {
		grpclog.Fatalln(err)
	}
	fmt.Println(r.Message)
}
