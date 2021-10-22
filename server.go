package main

import (
	"log"
	"net"

	"github.com/vulcan-dev/VKMP/chat"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatalf("failed to listen on port 1234: %v", err)
	}

	gs := grpc.NewServer()

	s := chat.Server{}
	chat.RegisterChatServiceServer(gs, &s)

	if err == gs.Serve(lis); err != nil {
		log.Fatalf("failed to serve gRPC: %v", err)
	}
}
