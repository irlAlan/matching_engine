package server

import (
	"log"
	"net"

	"github.com/irlalan/matching_engine/app/config"
	pb "github.com/irlalan/matching_engine/app/proto/order_service"

	"google.golang.org/grpc"
)

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
