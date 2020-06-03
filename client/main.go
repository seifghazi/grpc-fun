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
		resp, _ := client.SayHello(context.Background(), msg)
		log.Printf("Response from server: %s", resp)
	}
}

func main() {
	user := User{}
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("Please enter your name: ")
	for scanner.Scan() {
		name := scanner.Text()
		user.Name = name
		break
	}

	conn, err := grpc.Dial(":8000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %s", err)
	}
	defer conn.Close()

	client := proto.NewSendMessageClient(conn)

	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	go ReadAndSendInput(user, client)

	done := make(chan bool)
	<-done
}
