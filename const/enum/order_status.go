package enum

import (
	"go-gin/internal/etype"
)

// OrderStatus 订单状态
type OrderStatus struct {
	etype.BaseEnum
}

// 订单状态常量
var (
	ORDER_STATUS_PENDING   = etype.NewEnum[OrderStatus](1, "待支付")
	ORDER_STATUS_PAID      = etype.NewEnum[OrderStatus](2, "已支付")
	ORDER_STATUS_SHIPPING  = etype.NewEnum[OrderStatus](3, "配送中")
	ORDER_STATUS_COMPLETED = etype.NewEnum[OrderStatus](4, "已完成")
	ORDER_STATUS_CANCELLED = etype.NewEnum[OrderStatus](5, "已取消")
)
