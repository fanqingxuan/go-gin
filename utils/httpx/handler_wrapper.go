package httpx

import (
	"fmt"
	"go-gin/internal/errorx"
	"go-gin/svc"
	"reflect"

	"github.com/gin-gonic/gin"
)

type Handler interface {
	Init(ctx *gin.Context, svcCtx *svc.ServiceContext)
	Prepare()
	Before() error
	Handle(interface{}) (interface{}, error)
}

type BaseHandler struct {
	GinCtx *gin.Context
	*svc.ServiceContext
}

func (b *BaseHandler) Init(ginCtx *gin.Context, svcCtx *svc.ServiceContext) {
	b.GinCtx = ginCtx
	b.ServiceContext = svcCtx
}

func (b *BaseHandler) Before() error {
	return nil
}

func (b *BaseHandler) Prepare() {

}

func WrapHandler(h Handler, svcCtx *svc.ServiceContext, req interface{}) gin.HandlerFunc {
	var requestPtr interface{}
	var requestVal interface{}
	if req != nil {
		reqType := reflect.TypeOf(req)
		requestPtr = reflect.New(reqType).Interface()
	}
	return func(c *gin.Context) {
		if err := c.ShouldBind(requestPtr); err != nil {
			fmt.Println(err)
			Error(c, errorx.NewDefault(err.Error()))
			return
		}
		if req != nil {
			requestVal = reflect.ValueOf(requestPtr).Elem().Interface()
		}
		h.Init(c, svcCtx)
		h.Prepare()
		if err := h.Before(); err != nil {
			Error(c, err)
			return
		}
		resp, err := h.Handle(requestVal)
		if err != nil {
			Error(c, err)
			return
		}
		Ok(c, resp)
	}
}
