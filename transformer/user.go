package transformer

import (
	"go-gin/const/enum"
	"go-gin/internal/etype"
	"go-gin/model/entity"
	"go-gin/typing"
)

func ConvertUserToListData(u []*entity.User) []typing.ListData {
	var resp []typing.ListData
	for _, v := range u {
		var ageTips string
		var age int
		if v.Age != nil {
			age = *v.Age
			if *v.Age >= 18 {
				ageTips = "成年"
			} else {
				ageTips = "未成年"
			}
		}

		var statusText string
		if v.Status != nil {
			if status, err := etype.Parse[enum.UserStatus](*v.Status); err == nil {
				statusText = status.String()
			}
		}

		resp = append(resp, typing.ListData{
			Id:         int(v.Id),
			Name:       v.Name,
			AgeTips:    ageTips,
			Age:        age,
			Status:     v.Status,
			StatusText: statusText,
		})
	}
	return resp
}
