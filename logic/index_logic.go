package logic

import "context"

type IndexLogic struct {
}

func NewIndexLogic() *IndexLogic {
	return &IndexLogic{}
}

func (l *IndexLogic) Handle(ctx context.Context, req any) (resp any, err error) {
	return "user/index", nil
}
