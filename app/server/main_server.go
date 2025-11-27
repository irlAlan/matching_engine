package main // TODO: change this later

import (
	"context"
	"log"
	"net"

	pb "github.com/irlalan/matching_engine/app/proto/hello"

	// "google.golang.org/protobuf"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
	// hello.UninmplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Reply: "Hello " + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
