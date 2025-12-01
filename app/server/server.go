package server

import (
	ctx "context"
	"fmt"
	"log"
	"net"

	"github.com/irlalan/matching_engine/app/config"
	pb "github.com/irlalan/matching_engine/app/proto/order_service"

	// sv "github.com/irlalan/matching_engine/app/server"

	// "google.golang.org/protobuf"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedOrderServiceServer
}

func (s *server) CreateNewOrder(context ctx.Context, order *pb.Order) (*pb.Response, error) {
	Ob.Add_tobook(order)
	fmt.Println("Sending from the server: ")
	fmt.Printf("ob: %v\n", Ob)
	fmt.Println("Length: {}", len(Ob))

	return &pb.Response{StatusCode: 0, Response: "Completed"}, nil
}

// func (s *server) SayHello(ctx context.Context, in *hello_pb.HelloRequest) (*hello_pb.HelloReply, error) {
// 	return &hello_pb.HelloReply{Reply: "Hello " + in.Name}, nil
// }

func Start(conf *config.Config) {
	lis, err := net.Listen(conf.Connection.Network, conf.Connection.Ip_addr+conf.Connection.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	s.GetServiceInfo()
	pb.RegisterOrderServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
