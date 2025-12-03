package main

import (
	//	"github.com/irlalan/matching_engine/app/client"
	"github.com/irlalan/matching_engine/app/config"
	"github.com/irlalan/matching_engine/app/server"
)

// used to initialise everything

func main() {
	// load the config file for reading ip and port etc
	var conf config.Config
	conf.ReadConfig()
	server.Start(&conf)
	// client.Start(&conf)
}

//
//import (
//	"context"
//	"log"
//	"net"
//
//	hello_pb "github.com/irlalan/matching_engine/app/proto/hello"
//	pb "github.com/irlalan/matching_engine/app/proto/order_service"
//
//	// "google.golang.org/protobuf"
//	"google.golang.org/grpc"
//)
//
//type hello_server struct {
//	hello_pb.UnimplementedGreeterServer
//}
//
//type server struct {
//	pb.UnimplementedOrderServiceServer
//}
//
//func (s *hello_server) SayHello(ctx context.Context, in *hello_pb.HelloRequest) (*hello_pb.HelloReply, error) {
//	return &hello_pb.HelloReply{Reply: "Hello " + in.Name}, nil
//}
//
//func test() {
//	lis, err := net.Listen("tcp", ":50051")
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//	s := grpc.NewServer()
//	hello_pb.RegisterGreeterServer(s, &hello_server{})
//	// hello_pb.RegisterGreeterServer(s, &hello_server{})
//	log.Printf("server listening at %v", lis.Addr())
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//	// var ob OrderBook_tmp
//	// ob.Add_tobook(&pb.Order{})
//	// fmt.Println("Sending from the server: ")
//	// fmt.Printf("ob: %v\n", ob)
//}
//
