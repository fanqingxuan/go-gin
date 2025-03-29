package enum

import (
	"go-gin/internal/etype"
)

const PrefixOrderStatus etype.PrefixType = "order_status" // 订单状态前缀

// 定义订单状态常量
var (
	ORDER_STATUS_PENDING   = NewOrderStatus(1, "待支付")
	ORDER_STATUS_PAID      = NewOrderStatus(2, "已支付")
	ORDER_STATUS_SHIPPING  = NewOrderStatus(3, "配送中")
	ORDER_STATUS_COMPLETED = NewOrderStatus(4, "已完成")
	ORDER_STATUS_CANCELLED = NewOrderStatus(5, "已取消")
)

// OrderStatus 订单状态
type OrderStatus struct {
	etype.BaseEnum
}

// NewOrderStatus 创建订单状态
func NewOrderStatus(code int, desc string) *OrderStatus {
	return &OrderStatus{
		BaseEnum: etype.CreateBaseEnumAndSetMap(PrefixOrderStatus, code, desc),
	}
}

// ParseOrderStatus 解析订单状态
func ParseOrderStatus(code int) (*OrderStatus, error) {
	base, err := etype.ParseBaseEnum(PrefixOrderStatus, code)
	if err != nil {
		return nil, err
	}
	return &OrderStatus{BaseEnum: base}, nil
}

// Scan 实现 sql.Scanner 接口
func (s *OrderStatus) Scan(value interface{}) error {
	return s.BaseEnum.Scan(value, PrefixOrderStatus)
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *OrderStatus) UnmarshalJSON(data []byte) error {
	return s.BaseEnum.UnmarshalJSON(data, PrefixOrderStatus)
}
