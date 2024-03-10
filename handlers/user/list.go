package user

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/models"
	"go-gin/types"
	"go-gin/utils/httpx"
	"time"

	"github.com/golang-module/carbon/v2"
)

type ListUserHandler struct {
	httpx.BaseHandler
	userModel models.UserModel
}

func NewListUserHandler() *ListUserHandler {
	return new(ListUserHandler)
}

func (h *ListUserHandler) Prepare() {
	h.userModel = models.NewUserModel(h.GinCtx, h.DB)
}

func (h *ListUserHandler) Handle(request interface{}) (interface{}, error) {
	req, ok := request.(types.ListUserReq)
	fmt.Println(req, ok)

	if !ok {
		return nil, errorx.NewDefault("无效参数类型")
	}
	fmt.Println(carbon.Now().ToTimeString())

	h.Redis.Set(h.GinCtx, "tt", "dd", time.Hour)
	users, err := h.userModel.FindAll(3)
	if err != nil {
		return nil, err
	}

	return types.ListUserReply{
		User: users,
	}, nil
}
