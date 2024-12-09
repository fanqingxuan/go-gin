package logic

import "context"

type IndexLogic struct {
}

func NewIndexLogic() *IndexLogic {
	return &IndexLogic{}
}

func (l *IndexLogic) Handle(ctx context.Context, req interface{}) (resp interface{}, err error) {
	return "user/index", nil
}
