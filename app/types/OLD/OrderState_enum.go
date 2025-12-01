package types

type OrderStatus int

const (
	s_Open            OrderStatus = iota // 0
	s_PartiallyFilled                    // 1
	s_Filled                             // 2
	s_Cancelled                          // 3
)

var orderStatusStateName = map[OrderStatus]string{
	s_Open:            "open",
	s_PartiallyFilled: "paritally_filled",
	s_Filled:          "filled",
	s_Cancelled:       "cancelled",
}

func (ss OrderStatus) String() string {
	return orderStatusStateName[ss]
}
