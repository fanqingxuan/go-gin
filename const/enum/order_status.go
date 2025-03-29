package enum

import (
	"fmt"
	"go-gin/internal/etype"
)

const (
	PrefixOrderStatus etype.PrefixType = "order_status" // 订单状态前缀
)

// OrderStatus 订单状态
type OrderStatus struct {
	etype.BaseEnum
}

// 定义订单状态常量
var (
	ORDER_STATUS_PENDING   = NewOrderStatus(1, "待支付")
	ORDER_STATUS_PAID      = NewOrderStatus(2, "已支付")
	ORDER_STATUS_SHIPPING  = NewOrderStatus(3, "配送中")
	ORDER_STATUS_COMPLETED = NewOrderStatus(4, "已完成")
	ORDER_STATUS_CANCELLED = NewOrderStatus(5, "已取消")
)

// NewOrderStatus 创建订单状态
func NewOrderStatus(code int, desc string) *OrderStatus {
	status := &OrderStatus{
		BaseEnum: *etype.NewBaseEnum(code, desc),
	}
	etype.Set(PrefixOrderStatus, code, desc)
	return status
}

// ParseOrderStatus 解析订单状态
func ParseOrderStatus(code int) (*OrderStatus, error) {
	if desc, ok := etype.Get(PrefixOrderStatus, code); ok {
		return &OrderStatus{
			BaseEnum: *etype.NewBaseEnum(code, desc),
		}, nil
	}
	return nil, fmt.Errorf("未知的订单状态码: %d", code)
}

// Scan 实现 sql.Scanner 接口
func (s *OrderStatus) Scan(value interface{}) error {
	return s.BaseEnum.Scan(value, etype.GetAll(PrefixOrderStatus))
}

// UnmarshalJSON 实现 json.Unmarshaler 接口
func (s *OrderStatus) UnmarshalJSON(data []byte) error {
	return s.BaseEnum.UnmarshalJSON(data, etype.GetAll(PrefixOrderStatus))
}
