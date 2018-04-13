package main

import (
	"context"
	"testing"

	pb "github.com/gaku3601/Hello-Golang-gRPC/server/pb"
)

func TestGetHello(t *testing.T) {
	h := new(HelloService)
	ctx := context.Background()
	req := &pb.GetHelloRequest{Name: "gaku"}
	actual, _ := h.GetHello(ctx, req)
	expect := "こんにちわ！！gakuさん!"

	if actual.Greeting != expect {
		t.Errorf("\n結果:%#v\n予測:%#v", actual.Greeting, expect)
	}
}
