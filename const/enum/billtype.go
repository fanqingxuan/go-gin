package enum

import "fmt"

// 单据类型
type BillType int

const (
	// 正常添加
	BillTypeNormal BillType = 1
	// 从第三方导入
	BillTypeFromThird = 2
	// 退回生成的
	BillTypeReturn = 3
	// 调拨生成的
	BillTypeTransfer = 4
	// 盘点生成的
	BillTypeInventory = 5
)

var billTypeMap = map[BillType]string{
	BillTypeNormal:    "正常添加",
	BillTypeFromThird: "从第三方导入",
	BillTypeReturn:    "退回生成的",
	BillTypeTransfer:  "调拨生成的",
	BillTypeInventory: "盘点生成的",
}

func (b BillType) String() string {
	s, ok := billTypeMap[b]
	if !ok {
		return "未知"
	}
	return s
}

// Format 实现 fmt.Formatter 接口
func (b BillType) Format(f fmt.State, c rune) {
	// 对于所有的默认打印请求（如fmt.Println调用），使用text(xx)格式
	fmt.Fprintf(f, "%d", int(b))
}
