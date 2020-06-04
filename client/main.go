package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/grpc-fun/proto"
	"google.golang.org/grpc"
)

type User struct {
	Name string
}

func ReadAndSendInput(user User, client proto.SendMessageClient) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		msg := &proto.Message{
			Name:    user.Name,
			Message: scanner.Text(),
		}
		if _, err := client.SayHello(context.Background(), msg); err != nil {
			log.Fatalf("Error sending message: %s", err.Error())
		}
	}
}

func ReceiveMessage(stream proto.SendMessage_EstablishConnectionClient) {
	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message from client stream: %s", err.Error())
		}

		fmt.Printf("Message from %s: %s\n", msg.Name, msg.Message)
	}
}

func main() {
	fmt.Println("Attempting to connect to grpc server...")
	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
	}
	defer conn.Close()

	client := proto.NewSendMessageClient(conn)

	fmt.Println("Successfully connected!")
	fmt.Printf("Please enter your name: ")
	user := User{}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		name := scanner.Text()
		user.Name = name
		break
	}

	serverStream, err := client.EstablishConnection(context.Background(), &proto.Message{
		Name: user.Name,
	})
	if err != nil {
		log.Fatalf("Error establishing stream connection: %s", err.Error())
	}

	go ReceiveMessage(serverStream)
	go ReadAndSendInput(user, client)

	done := make(chan bool)
	<-done
}
