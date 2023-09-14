package server

import (
	"context"
	"log"
	"math"
	"voice-translator/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, req *proto.SqrtRequest) (resp *proto.SqrtResponse, err error) {
	log.Printf("Sqrt function was invoked with %v\n", req)

	if req.Num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, "Received a negative number: %v", req.Num)
	}

	return &proto.SqrtResponse{
		Ans: math.Sqrt(float64(req.Num)),
	}, nil
}
