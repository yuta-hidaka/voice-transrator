package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalcServiceClient) {
	log.Printf("start do sqrt ")
	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Num: -10})

	if err == nil {
		log.Printf("Sqrt: %v", res.Ans)
		return // success
	}

	e, ok := status.FromError(err)
	if !ok {
		log.Fatalf("failed to another reason besides grpc: %v", err)
		return
	}

	log.Printf("error code: %v", e.Code())
	log.Printf("error message: %v", e.Message())

	if e.Code() == codes.InvalidArgument {
		log.Printf("we probably sent a negative number!")
		return
	}

}

func TestCallingSqrt(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	doSqrt(c)
}
