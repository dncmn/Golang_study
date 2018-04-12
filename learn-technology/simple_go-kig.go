package main

import (
	"context"
	"errors"
	"strings"
	"github.com/go-kit/kit/endpoint"
	port "github.com/go-kit/kit/transport/http"
	"net/http"
	"encoding/json"
	"log"
)

// step-one:business logic

// StringService provides operations on strings
type StringService interface{
	Uppercase( context.Context, string)(string,error)
	Count(context.Context, string)int
}


// 实现接口


type stringService struct{}

// ErrEmpty is returned when input string is empty
var ErrEmpty =errors.New("Empty string")

func(_ stringService)Uppercase(_ context.Context, s string)(string,error){

	if s==""{
		return "",ErrEmpty
	}

	return  strings.ToUpper(s),nil
}

func(_ stringService)Count(_ context.Context,s string)int{
	return len(s)
}


// Requests and Responses
/*
	为每一个方法定义一个request and response 方便捕捉输入和输出参数
*/

type uppercaseRequest struct{
	S string `json:"s"`
}

type uppercaseResponse struct{
	V string `json:"v"`
	Err string `json:string,omitempty`
}

type countRequest struct {
	S string 	`json:"s"`
}

type countResponse struct{
	V int `json:"v"`
}

// Endpoints
/*
	一个Endpoint就是一个rpc,也就是我们定义的接口里面的一个方法,我们将采用
适配器来将我们定义在接口里面的方法转化成一个端点.每一个适配器将接收一个StringService.
并且返回住这个端点对应的那个接口里面的方法.
*/

func makeUppercaseEndpoint(svc StringService)endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req:=request.(uppercaseRequest)
		v,err:=svc.Uppercase(ctx,req.S)

		if err!=nil{
			return uppercaseResponse{v,err.Error()},nil
		}
		return uppercaseResponse{v,""},nil
	}
}



func makeCountEndpoint(svc StringService)endpoint.Endpoint{
	return func(ctx context.Context,request interface{})(interface{},error){
		req:=request.(countRequest)
		v:=svc.Count(ctx,req.S)
		return countResponse{v},nil
	}
}

// Transports

/*

	将我们的服务暴露在外界,以便他们调用。

	e endpoint.Endpoint,
	dec DecodeRequestFunc,
	enc EncodeResponseFunc,
	options ...ServerOption,
*/


func main(){
	svc:=stringService{}

	uppercaseHandler:=port.NewServer(makeUppercaseEndpoint(svc),decodeUppercaseRequest,encodeResponse)
	countHandler:=port.NewServer(makeCountEndpoint(svc),decodeCountRequest,encodeResponse)

	http.Handle("/uppercase",uppercaseHandler)
	http.Handle("/count",countHandler)

	log.Fatal(http.ListenAndServe(":8080",nil))
}




// decode and encode

func decodeUppercaseRequest(_ context.Context,r *http.Request)(interface{},error){
	var request uppercaseRequest

	if err:=json.NewDecoder(r.Body).Decode(&request);err!=nil{
		return nil ,err
	}

	return request,nil
}

func decodeCountRequest(_ context.Context,r *http.Request)(interface{},error){
	var request countRequest

	if err:=json.NewDecoder(r.Body).Decode(&request);err!=nil{
		return nil,err
	}

	return request,nil
}

func encodeResponse(_ context.Context,w http.ResponseWriter, response interface{})error{

	return json.NewEncoder(w).Encode(&response)

}












































