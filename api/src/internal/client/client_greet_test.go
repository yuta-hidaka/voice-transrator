package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGreet(c pb.GreetServiceClient) {
	log.Printf("start do greet")

	res, err := c.Greet(context.TODO(), &pb.GreetRequest{
		FirstName: "hehehehhehe",
	})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Greet: %s", res.FirstName)
}

func TestCallingGreet(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreet(c)
}
