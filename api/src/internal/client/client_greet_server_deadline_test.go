package client_test

import (
	"context"
	"log"
	"testing"
	"time"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func doGreetDeadLine(c pb.GreetServiceClient, timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetRequest{
		FirstName: "hehehehhehe",
	}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if !ok {
			log.Fatalf("failed to another reason besides grpc: %v", err)
		}

		if e.Code() == codes.DeadlineExceeded {
			log.Printf("timeout was hit! Deadline was exceeded")
			return
		}

		log.Printf("unexpected error: %v", e)

	}

	log.Printf("Greet: %s", res.FirstName)
}

func TestCallingDoGreetDeadLine(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreetDeadLine(c, 1*time.Second)
}
