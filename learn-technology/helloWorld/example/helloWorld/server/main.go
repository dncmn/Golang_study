package main

import (
	"golang.org/x/net/context"
	pb "helloWorld/example/proto"
	"fmt"
	"log"
	"net"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
)

const(
	Address ="127.0.0.1:50053"
)


// 定义weekDo接口
type weekDoService struct {}

var WeekDoService = weekDoService{}

// 实现定义的方法

func(w weekDoService)PlayGames(ctx context.Context, in *pb.WeekDoRequest)(*pb.WeekDoReply,error){
	resp:=new(pb.WeekDoReply)
	log.Println("sportsName=%s, sportsAge=%v",in.GetSportsName(),in.GetSportsAge())
	resp.ActivityNames=[]string{in.SportsName,fmt.Sprint(in.SportsAge,"")}
	return resp,nil
}

func(w weekDoService)WatchMoves(ctx context.Context,in *pb.WeekDoRequest)(*pb.WeekDoReply,error){
	resp:=new(pb.WeekDoReply)

	in.XTestCase_1="_test_case_1"
	in.TestCaseDemo_2="test_case_2"

	// 使用在.proto中定义的枚举类型
	log.Println(pb.SportsName_JAPAN_FOOTBALL)
	log.Println(pb.SportsName_CHINESE_BASKETBALL)
	log.Println(pb.SportsName_CHINESE_HAHA)
	resp.ActivityNames=[]string{in.SportsName}
	return resp,nil
}
// auth 验证Token
func auth(ctx context.Context) error {
	md, ok :=metadata.FromIncomingContext(ctx)

	if !ok {
		return grpc.Errorf(codes.Unauthenticated, "无Token认证信息")
	}

	var (
		appid  string
		appkey string
	)

	if val, ok := md["appid"]; ok {
		appid = val[0]
	}

	if val, ok := md["appkey"]; ok {
		appkey = val[0]
	}

	if appid != "101010" || appkey != "i am key" {
		return grpc.Errorf(codes.Unauthenticated, "Token认证信息无效: appid=%s, appkey=%s", appid, appkey)
	}

	return nil
}
func main(){
	listen,err:=net.Listen("tcp",Address)
	if err!=nil{
		log.Println("failed to listen:%v",err)
	}

	var opts []grpc.ServerOption

	// TLS认证
	creds, err := credentials.NewServerTLSFromFile("../../keys/server.pem", "../../keys/server.key")
	if err != nil {
		grpclog.Fatalf("Failed to generate credentials %v", err)
	}

	opts=append(opts,grpc.Creds(creds))
	// 注册interceptor
	var interceptor grpc.UnaryServerInterceptor
	
	interceptor= func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		err=auth(ctx)

		if err!= nil{
			return
		}
		// 继续处理请求
		return handler(ctx,req)
	}

	opts=append(opts,grpc.UnaryInterceptor(interceptor))
	
	
	// 实例化一个gRPC服务器
	server:=grpc.NewServer()

	// 注册一个WeekDoService服务
	pb.RegisterWeekDoServer(server,WeekDoService)

	log.Println("server is Listening :",Address)

	server.Serve(listen)

}















