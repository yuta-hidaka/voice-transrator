package client_test

import (
	"context"
	"io"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGreetStreamServer(c pb.GreetServiceClient) {
	log.Printf("start do greet stream server")

	req := &pb.GreetRequest{
		FirstName: "hehehehhehe",
	}

	stream, err := c.GreetServerStream(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("failed to receive: %v", err)
		}
		log.Printf("Greet: %s", res.FirstName)
	}
}

func TestCallingGreetStreamServer(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreetStreamServer(c)
}
