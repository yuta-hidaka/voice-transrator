package client_test

import (
	"context"
	"io"
	"log"
	"testing"
	"time"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGreetBiDirectionStream(c pb.GreetServiceClient) {
	log.Printf("start bi direction greet stream server")

	reqs := []*pb.GreetRequest{
		{
			FirstName: "hehehehhehe",
		},
		{
			FirstName: "hohohohohoho",
		},
		{
			FirstName: "hahahahahahah",
		},
	}
	stream, err := c.GreetBiDirectionStream(context.Background())
	if err != nil {
		log.Fatalf("failed to greet: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for _, req := range reqs {
			log.Print("sending request")
			stream.Send(req)
			time.Sleep(1 * time.Second)
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

			log.Printf("Greet: %s", res.FirstName)
		}
		close(waitc)
	}()

	<-waitc
}

func TestCallingDoGreetBiDirectionStream(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreetBiDirectionStream(c)
}
