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
	connections []*Connection
}

type Connection struct {
	Id     string
	Name   string
	Stream proto.SendMessage_ServerStreamHelloServer
	Error  chan error
}

// SayHello generates response to a Message request
func (s *Server) SayHello(ctx context.Context, m *proto.Message) (*proto.Message, error) {
	for _, conn := range s.connections {
		if m.Name != conn.Name {
			conn.Stream.Send(m)
		}
	}
	return nil, nil
}

func (s *Server) EstablishConnection(message *proto.Message, stream proto.SendMessage_EstablishConnectionServer) error {
	newConnection := &Connection{
		Id:     uuid.New().String(),
		Name:   message.Name,
		Stream: stream,
		Error:  make(chan error),
	}

	s.connections = append(s.connections, newConnection)

	msg := proto.Message{
		Name:    "Server",
		Message: fmt.Sprintf("Welcome to the chat %s! There are %d other clients connected", newConnection.Name, len(s.connections)-1),
	}
	stream.Send(&msg)

	return <-newConnection.Error
}

func (s *Server) ServerStreamHello(msg *proto.Message, stream proto.SendMessage_ServerStreamHelloServer) error {
	for _, conn := range s.connections {
		if msg.Name != conn.Name {
			conn.Stream.Send(msg)
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
