package types

import (
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type Order struct {
	id     int
	side   OrderSide
	price  int
	volume int
	o_type OrderType
	status OrderStatus
	name   string
	time   timestamppb.Timestamp
}
