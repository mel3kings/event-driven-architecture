package main

import (
	"context"
	pb "github.com/mel3kings/event-driven-architecture/grpc/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("greet function invoked with params %v", in)
	return &pb.GreetResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
