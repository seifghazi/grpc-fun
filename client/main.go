package main

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/grpc-fun/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
	}
	defer conn.Close()

	client := proto.NewSendMessageClient(conn)

	resp, err := client.SayHello(context.Background(), &proto.Message{Id: uuid.New().String(), Message: "Client m8"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", resp)
}
