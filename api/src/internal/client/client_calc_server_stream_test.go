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

func doCalcStreamServer(c pb.CalcServiceClient) {
	log.Printf("start do Calc stream server")

	req := &pb.CalcCalcServerStreamRequest{
		Num: 120,
	}

	stream, err := c.CalcServerStream(context.Background(), req)
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

		log.Printf("Calc: %v", res.Ans)
	}
}

func TestCallingCalcStreamServer(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	doCalcStreamServer(c)
}
