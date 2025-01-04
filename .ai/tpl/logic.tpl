package logic

import (
	"context"
	"fmt"
	"go-gin/model"
	"go-gin/typing"
)

type XxxLogic struct {
	model *model.模型名称
}

func NewXxxLogic() *XxxLogic {
	return &XxxLogic{
		model: model.New模型名称(),
	}
}

func (l *XxxLogic) Handle(ctx context.Context, req typing.xxxReq) (resp *typing.xxxResp, err error) {
	resp = &typing.xxxResp{
	}
	return
}
