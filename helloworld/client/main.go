package main

import (
	"context"
	"fmt"
	"github.com/kaikaichao/grpctest/pb/hello"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

)

func main() {
	opts:=grpc.WithTransportCredentials(insecure.NewCredentials()) //设置安全属性（设置为不安全的）
	conn,err:=grpc.Dial("localhost:8888",opts)
	if err!=nil{
		fmt.Println(err)
	}
	client:=hello.NewHelloGRPCClient(conn)
	res,_:=client.SayHi(context.Background(),&hello.Req{Name: "我从客户端来"})
	fmt.Println(res.Message,"message")
	fmt.Println(res.GetMessage(),"getmessage")
}
