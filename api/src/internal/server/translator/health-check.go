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

func (s *Server) HealthCheckStream(stream pb.TranslatorService_HealthCheckStreamServer) error {
	times := 0
	for {
		_, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatal(err)
			return err
		}

		err = stream.Send(&pb.HealthCheckResponse{
			Health: fmt.Sprintf("hello %v times of response", times),
		})
		if err != nil {
			log.Fatal(err)
			return err
		}
	}
}
