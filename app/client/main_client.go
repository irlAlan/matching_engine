package main // TODO: change this later

import (
	"context"
	"log"
	"time"

	hello_pb "github.com/irlalan/matching_engine/app/proto/hello"
	pb "github.com/irlalan/matching_engine/app/proto/order_service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/timestamppb"
	// hello_pb "google.golang.org/protobuf"
)

func main() {

	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //, grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	c := pb.NewOrderServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.CreateNewOrder(ctx, &pb.Order{
		Id:     &pb.OrderID{Id: "test"},
		Name:   "Alan Test",
		Side:   pb.Side_BUY,
		Price:  12345,
		Volume: 1,
		Type:   pb.Type_LIMIT_ORDER,
		Status: pb.Status_OPEN,
		Time:   timestamppb.Now(),
	})

	v, err := c.CreateNewOrder(ctx, &pb.Order{
		Id:     &pb.OrderID{Id: "test2"},
		Name:   "Alan Test",
		Side:   pb.Side_BUY,
		Price:  12345,
		Volume: 1,
		Type:   pb.Type_LIMIT_ORDER,
		Status: pb.Status_OPEN,
		Time:   timestamppb.Now(),
	})

	if err != nil {
		log.Fatalf("could not create new Order: %v", err)
	}
	log.Printf("Response: %s", r.GetResponse())
	log.Printf("Response: %s", v.GetResponse())
}

func hello() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials())) //, grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := hello_pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &hello_pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetReply())
}
