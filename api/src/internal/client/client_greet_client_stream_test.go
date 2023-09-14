package client_test

import (
	"context"
	"log"
	"testing"

	pb "voice-translator/internal/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func doGreetClientStream(c pb.GreetServiceClient) {
	log.Printf("start do greet client stream")

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
	stream, err := c.GreetClientStream(context.Background())
	if err != nil {
		log.Fatalf("failed to while calling greet: %v", err)
	}

	for _, req := range reqs {
		log.Print("sending request")
		stream.Send(req)
		// time.Sleep( * time.Second)
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to receive: %v", err)
	}
	log.Printf("Greet: %s", res.FirstName)
}

func TestCallingDoGreetClientStream(m *testing.T) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	doGreetClientStream(c)
}
