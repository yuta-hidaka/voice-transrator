package server

import (
	"context"
	"log"
	"time"
	pb "voice-translator/internal/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, req *pb.GreetRequest) (resp *pb.GreetResponse, err error) {
	log.Printf("GreetWithDeadline function was invoked with %v\n", req)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request!")
		}

		time.Sleep(1 * time.Second)
	}

	resp = &pb.GreetResponse{
		FirstName: req.FirstName,
	}

	return
}
