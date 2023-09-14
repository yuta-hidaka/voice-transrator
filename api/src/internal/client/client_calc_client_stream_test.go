package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doCalcClientStream(c pb.CalcServiceClient) {
	log.Printf("start do Calc stream client")

	stream, err := c.CalcClientStream(context.Background())
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	reqs := []*pb.CalcCalcServerStreamRequest{
		{
			Num: 1,
		},
		{
			Num: 2,
		},
		{
			Num: 3,
		},
		{
			Num: 4,
		},
	}
	for _, v := range reqs {
		stream.Send(v)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}
	log.Printf("Calc ave: %v", res.Ans)
}

func TestCallingDoCalcClientStream(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	doCalcClientStream(c)
}
