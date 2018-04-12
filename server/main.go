package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/gaku3601/Hello-Golang-gRPC/server/pb"

	"google.golang.org/grpc"
)

func main() {
	listenPort, err := net.Listen("tcp", ":19003")
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	s := &HelloService{}
	// 実行したい実処理をseverに登録する
	pb.RegisterHelloServer(server, s)
	server.Serve(listenPort)
}

type HelloService struct{}

func (h *HelloService) GetHello(ctx context.Context, req *pb.GetHelloRequest) (*pb.GetHelloResponse, error) {
	return &pb.GetHelloResponse{Greeting: fmt.Sprintf("こんにちわ！！%vさん!", req.Name)}, nil
}
