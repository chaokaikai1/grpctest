package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	hello_world "github.com/kaikaichao/grpctest/pb/gateway"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"net"
	"net/http"
	"sync"
)

type helloServer struct {
	hello_world.UnimplementedGreeterServer
}
func (*helloServer) SayHello(ctx context.Context,req *hello_world.HelloRequest) (*hello_world.HelloRes, error) {
	name:=req.Name
	res:=&hello_world.HelloRes{Msg: fmt.Sprintf("my name is %s",name)}
	return res,nil
}
func main() {
	wg:=&sync.WaitGroup{}
	wg.Add(2)
	go registGateWay(wg)
	go registGRPC(wg)
	wg.Wait()

}

func registGRPC(wg *sync.WaitGroup){
	defer  wg.Done()
	l,_:=net.Listen("tcp",":8888")
	s:=grpc.NewServer()
	hello_world.RegisterGreeterServer(s,&helloServer{})
	s.Serve(l)
}

func  registGateWay(wg *sync.WaitGroup){
	defer  wg.Done()
	opts:=grpc.WithTransportCredentials(insecure.NewCredentials())
	conn,_:=grpc.Dial("127.0.0.1:8888",opts)
	mux:=runtime.NewServeMux()
	gwServer:=&http.Server{
		Handler: mux,
		Addr: ":8090",
	}
	_=hello_world.RegisterGreeterHandler(context.Background(),mux,conn)
	gwServer.ListenAndServe()
}
