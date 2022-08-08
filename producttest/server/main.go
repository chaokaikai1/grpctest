package main

import (
	"context"
	"fmt"
	"github.com/kaikaichao/grpctest/pb/product"
	"google.golang.org/grpc"
	"net"
	"time"
)

type productServer struct {
	product.UnimplementedSearchServiceServer
}
func (*productServer) Search(ctx context.Context, req *product.ProductReq) (*product.ProductRes, error) {
	name:=req.GetName()
	s:=fmt.Sprintf("my productName is %s",name)
	res:=&product.ProductRes{Name: s}
	return res,nil
}
func (*productServer) SearchIn(server product.SearchService_SearchInServer) error {
	for{
		req,err:=server.Recv()
		fmt.Println(req)
		if err!=nil{
			fmt.Println(err)
			server.SendAndClose(&product.ProductRes{Name: "接收结束了"})
			break
		}
	}
	return nil
}
func (*productServer) SearchOut(req *product.ProductReq, server product.SearchService_SearchOutServer) error {
	name:=req.Name
	for i:=0;i<10;i++{
		time.Sleep(1*time.Second)
		newName:=fmt.Sprintf("我是out的名称 %s,%d",name,i)
		server.SendMsg(&product.ProductRes{Name: newName})
	}

	return nil
}
func (*productServer) SearchIO(product.SearchService_SearchIOServer) error {
	return nil
}

func main() {
	l,_:=net.Listen("tcp",":8888")
	s:=grpc.NewServer()
	product.RegisterSearchServiceServer(s,&productServer{})
	s.Serve(l)
}
