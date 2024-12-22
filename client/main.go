package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/Davinder1436/Go_GRPC/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)

	// Unary Call
	unaryResponse, err := client.UnaryCall(context.Background(), &pb.Input{Message: "Hello"})
	if err != nil {
		log.Fatalf("Unary call failed: %v", err)
	}
	log.Println("Unary response:", unaryResponse.Response)

	// Client Streaming Call
	clientStream, err := client.ClientStreamingCall(context.Background())
	if err != nil {
		log.Fatalf("Client streaming call failed: %v", err)
	}
	for _, msg := range []string{"Hello", "from", "dave", "recieved", "success"} {
		if err := clientStream.Send(&pb.Input{Message: msg}); err != nil {
			log.Fatalf("Failed to send: %v", err)
		}
	}
	clientStreamResponse, err := clientStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to receive response: %v", err)
	}
	log.Println("Client streaming response:", clientStreamResponse.Response)

	// Server Streaming Call
	serverStream, err := client.ServerStreamingCall(context.Background(), &pb.Input{Message: "Hello"})
	if err != nil {
		log.Fatalf("Server streaming call failed: %v", err)
	}
	for {
		resp, err := serverStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		log.Println("Server streaming response:", resp.Response)
	}

	// Bidirectional Streaming Call
	bidirectionalStream, err := client.BidirectionalStreamingCall(context.Background())
	if err != nil {
		log.Fatalf("Bidirectional streaming call failed: %v", err)
	}
	go func() {
		for _, msg := range []string{"Hello", "from", "dave", "recieved", "success"} {
			if err := bidirectionalStream.Send(&pb.Input{Message: msg}); err != nil {
				log.Fatalf("Failed to send: %v", err)
			}
			time.Sleep(time.Second)
		}
		bidirectionalStream.CloseSend()
	}()
	for {
		resp, err := bidirectionalStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive: %v", err)
		}
		log.Println("Bidirectional streaming response:", resp.Response)
	}
}
