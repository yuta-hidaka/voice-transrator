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

func doCalcBiDirectionStreamServer(c pb.CalcServiceClient) {
	log.Printf("start bi direction Calc stream server")

	// 1,5,3,6,2,20
	reqs := []*pb.CalcCalcServerStreamRequest{
		{
			Num: 1,
		},
		{
			Num: 5,
		},
		{
			Num: 3,
		},
		{
			Num: 6,
		},
		{
			Num: 2,
		},
		{
			Num: 20,
		},
	}

	stream, err := c.CalcBiDirectionStream(context.Background())
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Print("sending request")
			stream.Send(req)
		}
		stream.CloseSend()
	}()

	go func() {
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

		close(waitc)
	}()
	<-waitc
}

func TestCallingDoCalcBiDirectionStreamServer(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewCalcServiceClient(conn)

	doCalcBiDirectionStreamServer(c)
}
