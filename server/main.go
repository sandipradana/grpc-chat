package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/sandipradana/grpc-chat/model"
	"google.golang.org/grpc"
)

type service struct {
	pb.UnimplementedChatServiceServer
}

func (s *service) Send(ctx context.Context, in *pb.Message) (*pb.Message, error) {
	fmt.Println("Client : ", in.GetSender(), " Message : ", in.GetBody())
	return &pb.Message{Sender: "Server", Body: in.GetSender() + ", your message is " + in.GetBody()}, nil
}

func main() {
	tcpServer, err := net.Listen("tcp", ":9000")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, &service{})

	if err := grpcServer.Serve(tcpServer); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
