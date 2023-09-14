package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doCalc(c pb.CalcServiceClient) {
	log.Printf("start do greet")

	res, err := c.Calc(context.TODO(), &pb.CalcRequest{
		One: 1,
		Two: 2,
	})
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}
	log.Printf("Calc: %v", res.Ans)
}

func TestCallingCalc(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	doCalc(c)
}
