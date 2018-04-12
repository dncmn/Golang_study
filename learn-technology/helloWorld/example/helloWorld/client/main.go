package main

import (
	pb "helloWorld/example/proto"
	"google.golang.org/grpc"
	"log"
	"context"
)

// 设置监听的服务器的地址
const(
	Address ="127.0.0.1:50053"
)

func main(){
	conn,err:=grpc.Dial(Address,grpc.WithInsecure())
	if err!=nil{
		log.Fatalln(err)
	}

	defer conn.Close()

	// 初始化一个客户端
	client:=pb.NewWeekDoClient(conn)

	// 调用方法
	reqBody:=new(pb.WeekDoRequest)
	reqBody.SportsName="basketball"
	reqBody.SportsAge=18
	reqBody.SportsKinds=[]string{"English","math","pe"}

	res,err:=client.PlayGames(context.Background(),reqBody)
	if err!=nil{
		log.Fatalln(err)
	}

	log.Println(res.ActivityNames)
}
