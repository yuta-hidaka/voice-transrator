package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGreet(c pb.TranslatorServiceClient) {
	log.Printf("start do health-check")
	res, err := c.HealthCheck(context.TODO(), &pb.HealthCheckRequest{})
	if err != nil {
		log.Fatalf("failed to health-check: %v", err)
	}
	log.Printf("Greet: %s", res.Health)
}

func TestCallingGreet(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewTranslatorServiceClient(conn)

	doGreet(c)
}
