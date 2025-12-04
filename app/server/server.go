package server

import (
	ctx "context"
	"fmt"

	pb "github.com/irlalan/matching_engine/app/proto/order_service"
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
