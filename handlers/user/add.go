package user

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/logic/user"
	"go-gin/svc"
	"go-gin/types"
	"go-gin/utils/httpx"
	"time"

	"github.com/gin-gonic/gin"
)

func AddUser(serverCtx *svc.ServiceContext) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req types.AddUserReq
		if err := ctx.ShouldBind(&req); err != nil {
			httpx.Error(ctx, errorx.NewDefault(err.Error()))
			return
		}

		l := user.NewAddUser()
		resp, err := l.Handle(req)
		if err != nil {
			httpx.Error(ctx, err)
			return
		}
		result := serverCtx.Redis.Set(ctx, "name", "测试", time.Duration(time.Hour))
		fmt.Println(result.Err())
		fmt.Print(serverCtx.Redis.Get(ctx, "name"))
		httpx.Ok(ctx, resp)

	}
}
