package transformer

import (
	"go-gin/model"
	"go-gin/typing"
)

func ConvertUserToListData(u []*model.User) []typing.ListData {
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

		resp = append(resp, typing.ListData{
			Id:           int(v.Id),
			Name:         v.Name,
			AgeTips:      ageTips,
			Age:          age,
			Status:       v.Status,
			UserType:     v.UserType,
			UserTypeText: v.Status.String(),
		})
	}
	return resp
}
