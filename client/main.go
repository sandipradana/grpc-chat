package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/sandipradana/grpc-chat/model"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Chat Client")
	fmt.Println("--------------------")

	fmt.Print("Name : ")
	sender, _ := reader.ReadString('\n')
	sender = sender[:len(sender)-2]

	for {

		fmt.Print("Message : ")
		message, _ := reader.ReadString('\n')
		fmt.Println("")

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)

		c := pb.NewChatServiceClient(conn)
		r, err := c.Send(ctx, &pb.Message{Sender: sender, Body: message})

		if err != nil {
			log.Fatalf("Error: %v", err)
		} else {
			fmt.Println("Server : ", r.GetBody())
		}

		cancel()
	}
}
