package main

import (
	"context"
	pb "github.com/mel3kings/event-driven-architecture/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect : %v", err)
	}
	log.Println("connected")

	c := pb.NewGreetServiceClient(conn)
	doGreet(c)
	defer conn.Close()
}

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	r, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Mel"})

	if err != nil {
		log.Fatalf("Could not greet: %v\n", err)
	}

	log.Print(r.Result)
}
