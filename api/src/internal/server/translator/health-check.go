package translator

import (
	"context"
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
