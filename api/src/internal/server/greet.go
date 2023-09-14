package server

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "voice-translator/internal/proto"
)

func (s *Server) Greet(ctx context.Context, req *pb.GreetRequest) (*pb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with %v\n", req)
	firstName := req.FirstName
	result := "Hello " + firstName
	res := &pb.GreetResponse{
		FirstName: result,
	}
	return res, nil
}

func (s *Server) GreetServerStream(req *pb.GreetRequest, stream pb.GreetService_GreetServerStreamServer) error {
	fmt.Printf("Greet many times function was invoked with %v\n", stream)

	for i := 0; i < 10; i++ {
		result := "Hello " + req.FirstName + " number " + fmt.Sprint(i)
		stream.Send(&pb.GreetResponse{
			FirstName: result,
		})
		fmt.Printf("Greet many times function was invoked with the result %v\n", result)
	}

	return nil
}

func (s *Server) GreetClientStream(stream pb.GreetService_GreetClientStreamServer) error {
	log.Println("GreetClientStream function was invoked with a streaming request")
	res := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				FirstName: res,
			})
		}
		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}
		res += fmt.Sprintf("Hello %s! ", req.FirstName)
	}
}

func (s *Server) GreetBiDirectionStream(stream pb.GreetService_GreetBiDirectionStreamServer) error {
	log.Println("Greet bi direction stream function was invoked with a streaming request")
	for {

		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if err := stream.Send(&pb.GreetResponse{
			FirstName: fmt.Sprintf("Hello %s! ", req.FirstName),
		}); err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}

	}
}
