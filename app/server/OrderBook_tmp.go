package server

import (
	order_service "github.com/irlalan/matching_engine/app/proto/order_service"
)

type orderBook_tmp map[*order_service.OrderID]*order_service.Order

var Ob = make(orderBook_tmp)

//	type Something struct {
//		ob orderBook_tmp
//	}
//
//	func Init() *Something {
//		return &Something{ob: make(orderBook_tmp)}
//	}
func (ss *orderBook_tmp) Add_tobook(os *order_service.Order) string {
	Ob[os.Id] = os
	return "doing something"
}
