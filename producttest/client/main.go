package main

import (
	"context"
	"fmt"
	"github.com/kaikaichao/grpctest/pb/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
		opts:=grpc.WithTransportCredentials(insecure.NewCredentials())
		conn,_:=grpc.Dial("127.0.0.1:8888",opts)
		client:=product.NewSearchServiceClient(conn)
		//普通的调用
		//res,_:=client.Search(context.Background(),&product.ProductReq{Name: "三只松鼠"})
		//fmt.Println(res)

		//流式输入
		//serverClient,_:=client.SearchIn(context.Background())
		//for i:=0;i<5;i++{
		//	time.Sleep(1*time.Second)
		//	serverClient.Send(&product.ProductReq{Name: "流式传入:"+strconv.Itoa(i)})
		//}
		//res,_:=serverClient.CloseAndRecv()
		//fmt.Println(res)

		//流式输出
		c,_:=client.SearchOut(context.Background(),&product.ProductReq{Name:"ckk"})
		for {
			res,err:=c.Recv()
			if err!=nil{
				fmt.Println(err)
				break
			}
			fmt.Println(res)
		}

}
