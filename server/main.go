package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/Davinder1436/Go_GRPC/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedExampleServiceServer
}

func (s *server) UnaryCall(ctx context.Context, req *pb.Input) (*pb.Output, error) {
	return &pb.Output{Response: "Received: " + req.Message}, nil
}

func (s *server) ClientStreamingCall(stream pb.ExampleService_ClientStreamingCallServer) error {
	var combinedMessage string
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			// Send the final response once all input has been received.
			return stream.SendAndClose(&pb.Output{Response: combinedMessage})
		}
		if err != nil {
			return err
		}
		combinedMessage += req.Message + " "
	}
}

func (s *server) ServerStreamingCall(req *pb.Input, stream pb.ExampleService_ServerStreamingCallServer) error {
	for i := 0; i < 5; i++ {
		response := fmt.Sprintf("Stream %d: Received %s", i, req.Message)
		if err := stream.Send(&pb.Output{Response: response}); err != nil {
			return err
		}
	}
	return nil
}

func (s *server) BidirectionalStreamingCall(stream pb.ExampleService_BidirectionalStreamingCallServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		response := "Echo: " + req.Message
		if err := stream.Send(&pb.Output{Response: response}); err != nil {
			return err
		}
	}
}

func main() {
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterExampleServiceServer(grpcServer, &server{})
	log.Println("Server is running on port 50051")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
