package main

import (
	"context"
	pb "github.com/mel3kings/event-driven-architecture/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
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
	doGreatMany(c)
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

func doGreatMany(c pb.GreetServiceClient) {
	stream, err := c.MultipleGreat(context.Background(), &pb.GreetRequest{FirstName: "Multiple"})
	if err != nil {
		log.Fatalf("error setting up connection %v", err)
	}

	for {
		msg, streamErr := stream.Recv()
		if streamErr == io.EOF {
			break
		}
		if streamErr != nil {
			log.Fatalf("error setting up connection %v", err)
		}
		log.Printf("%s", msg.Result)
	}
}
