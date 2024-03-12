package user

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/models"
	"go-gin/types"
	"go-gin/utils/httpx"
	"go-gin/validators"
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

type T struct {
	Name string `binding:"required" label:"学生姓名"`
}

func (h *ListUserHandler) Handle(request interface{}) (interface{}, error) {
	req, ok := request.(types.ListUserReq)
	fmt.Println(req)
	if !ok {
		return nil, errorx.NewDefault("无效参数类型")
	}
	fmt.Println(carbon.Now().ToTimeString())
	t := T{""}
	if err := validators.Validate(t); err != nil {
		return nil, err
	}
	h.Redis.Set(h.GinCtx, "tt", "dd", time.Hour)
	users, err := h.userModel.FindAll(3)
	if err != nil {
		return nil, err
	}

	return types.ListUserReply{
		User: users,
	}, nil
}
