package main

import (
	"context"
	"fmt"
	"github.com/kaikaichao/grpctest/pb/hello"
	"google.golang.org/grpc"
	"net"
)

type server struct {
	hello.UnimplementedHelloGRPCServer
}
func (s *server) SayHi(ctx context.Context, req *hello.Req) ( res *hello.Res, err error) {
	fmt.Println(req.Name)
	return  &hello.Res{Message: "我是从服务端返回的grpc内容"},nil
}
func main() {
	//注册服务
	l,_:=net.Listen("tcp",":8888")
	s:=grpc.NewServer()
	hello.RegisterHelloGRPCServer(s,&server{})
	//建立监听
	s.Serve(l)
}
