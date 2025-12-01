package types

type OrderType int

const (
	s_LimitOrder  OrderType = iota // only at specific price or better
	s_MarketOrder                  // Order immediatly at the best price
	s_StopOrder                    // conditional that become one of the others once a specific price is reached
)

var orderTypeStateName = map[OrderType]string{
	s_LimitOrder:  "limit_order",
	s_MarketOrder: "market_order",
	s_StopOrder:   "stop_order",
}

func (ss OrderType) String() string {
	return orderTypeStateName[ss]
}
