package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/grpc-fun/proto"
	"google.golang.org/grpc"
)

type Server struct {
	connections []Connection
}

type Connection struct {
	Name   string
	stream proto.SendMessage_ServerStreamHelloServer
	error  chan error
}

// SayHello generates response to a Message request
func (s *Server) SayHello(ctx context.Context, m *proto.Message) (*proto.Message, error) {
	fmt.Printf("Received message from %s: %s", m.Name, m.Message)
	Name := uuid.New().String()
	return &proto.Message{Name: Name, Message: "bar"}, nil
}

func (s *Server) ServerStreamHello(msg *proto.Message, stream proto.SendMessage_ServerStreamHelloServer) error {
	conn := Connection{
		Name:   msg.Name,
		stream: stream,
		error:  make(chan error),
	}

	s.connections = append(s.connections, conn)

	for i := 0; i < 10; i++ {
		err := stream.Send(&proto.Message{
			Name:    uuid.New().String(),
			Message: fmt.Sprintf("Hellooo client: %s", msg.Name),
		})

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Server) ClientStreamHello(clientStream proto.SendMessage_ClientStreamHelloServer) error {
	for {
		msg, err := clientStream.Recv()

		if err == io.EOF {
			return clientStream.SendAndClose(&proto.Message{
				Name:    uuid.New().String(),
				Message: "Buh Bye m8",
			})
		}

		fmt.Printf("Message from client: %s\n", msg)
	}
}

func (s *Server) BiDricetionalStreamHello(stream proto.SendMessage_BiDricetionalStreamHelloServer) error {
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		stream.Send(&proto.Message{
			Name:    uuid.New().String(),
			Message: fmt.Sprintf("Got your message: %s", msg.Message),
		})
	}
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := &Server{}

	grpcServer := grpc.NewServer()

	proto.RegisterSendMessageServer(grpcServer, s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
