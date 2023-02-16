package main

import (
	"context"
	"fmt"
	pb "github.com/mel3kings/event-driven-architecture/grpc/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function invoked with params %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}

func (s *Server) MultipleGreat(in *pb.GreetRequest, stream pb.GreetService_MultipleGreatServer) error {
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s %d", in.FirstName, i)
		stream.Send(&pb.GreetResponse{Result: res})
	}

	return nil
}
