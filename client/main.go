package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/gaku3601/Hello-Golang-gRPC/client/pb"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:19003", grpc.WithInsecure())
	if err != nil {
		log.Fatal("client connection error:", err)
	}
	defer conn.Close()
	client := pb.NewHelloClient(conn)
	message := &pb.GetHelloRequest{Name: "gaku( ´ ▽ ` )"}
	res, err := client.GetHello(context.TODO(), message)
	fmt.Printf("%v \n", res.Greeting)
	fmt.Printf("error::%#v \n", err)

	message2 := &pb.GetProfileRequest{Name: "gaku( ´ ▽ ` )", Age: 27}
	res2, err := client.GetProfile(context.TODO(), message2)
	fmt.Printf("%v \n", res2.Info)
	fmt.Printf("error::%#v \n", err)
}
