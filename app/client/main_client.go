package main // TODO: change this later

import (
	"context"
	"log"
	"time"

	pb "github.com/irlalan/matching_engine/app/proto/hello"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	// pb "google.golang.org/protobuf"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //, grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
