package enum

var (
	STATUS_VALID   = NewStatus(1, "有效")
	STATUS_INVALID = NewStatus(2, "无效")
	STATUS_DELETED = NewStatus(3, "已删除")
)
