package server

import (
	"context"
	"fmt"
	"io"
	"log"
	pb "voice-translator/internal/proto"
)

type Server struct {
	pb.GreetServiceServer
	pb.CalcServiceServer
}

func (s *Server) Calc(ctx context.Context, req *pb.CalcRequest) (*pb.CalcResponse, error) {
	fmt.Printf("Calc function was invoked with %v\n", req)
	return &pb.CalcResponse{
		Ans: req.One + req.Two,
	}, nil
}

func (s *Server) CalcServerStream(req *pb.CalcCalcServerStreamRequest, stream pb.CalcService_CalcServerStreamServer) error {
	var prime int64 = 2
	value := req.Num

	for value > 1 {
		if value%prime == 0 {
			stream.Send(&pb.CalcCalcServerStreamResponse{
				Ans: prime,
			})
			value /= prime
			continue
		}
		prime++
	}

	return nil
}

func (s *Server) CalcClientStream(stream pb.CalcService_CalcClientStreamServer) error {

	ans := int64(0)
	count := int64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.CalcCalcClientStreamResponse{
				Ans: float32(ans) / float32(count),
			})
		}
		if err != nil {
			log.Fatalln(err)
		}
		ans += req.Num
		count++
	}
}

func (s *Server) CalcBiDirectionStream(stream pb.CalcService_CalcBiDirectionStreamServer) error {
	log.Println("Calc bi direction stream function was invoked with a streaming request")

	prev := int64(0)

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		if prev > req.Num {
			continue
		}

		prev = req.Num

		if err := stream.Send(&pb.CalcCalcServerStreamResponse{
			Ans: req.Num,
		}); err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
	}

}
