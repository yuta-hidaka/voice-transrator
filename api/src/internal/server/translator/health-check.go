package translator

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "voice-translator/internal/proto"
)

type Server struct {
	pb.TranslatorServiceServer
}

func (s *Server) HealthCheck(_ context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{
		Health: "OK",
	}, nil
}

func (s *Server) HealthCheckClientStream(stream pb.TranslatorService_HealthCheckClientStreamServer) error {
	cnt := 0
	fmt.Println("hello111111")
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			fmt.Println("eof")
			return nil
		}

		fmt.Println("hello")

		if err != nil {
			fmt.Println("errors")
			log.Fatal(err)
			return err
		}

		cnt += 1
		fmt.Println("hello")
	}

	return nil
}

func (s *Server) HealthCheckBiDiStream(stream pb.TranslatorService_HealthCheckBiDiStreamServer) error {
	cnt := 0
	for {
		stream.Send(&pb.HealthCheckResponse{
			Health: fmt.Sprintf("OK %d", cnt),
		})

		v, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		fmt.Println("my hello is ", v.Hello)

		if err != nil {
			log.Fatal(err)
			return err
		}
		cnt += 1
	}

	return nil
}
