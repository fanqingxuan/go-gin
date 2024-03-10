package user

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/types"
	"go-gin/utils/httpx"
	"time"
)

type ListUserHandler struct {
	httpx.BaseHandler
}

func NewListUserHandler() *ListUserHandler {
	return new(ListUserHandler)
}

func (h *ListUserHandler) Handle(request interface{}) (interface{}, error) {
	req, ok := request.(types.ListUserReq)
	fmt.Println(req, ok)

	if !ok {
		return nil, errorx.NewDefault("无效参数类型")
	}
	h.Redis.Set(h.GinCtx, "tt", "dd", time.Hour)
	// return nil, errorx.NewDefault("查询数据不存在")
	return types.ListUserReply{
		User: []types.User{
			{
				Name: "测试",
			},
			{
				Name: "测试222",
			},
		},
	}, nil
}
