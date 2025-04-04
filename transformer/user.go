package transformer

import (
	"go-gin/model"
	"go-gin/typing"
)

func ConvertUserToListData(u []model.User) []typing.ListData {
	var resp []typing.ListData
	for _, v := range u {
		var ageTips string
		if v.Age != nil && *v.Age >= 18 {
			ageTips = "成年"
		} else {
			ageTips = "未成年"
		}
		resp = append(resp, typing.ListData{
			Id:            int(v.Id),
			Name:          v.Name,
			AgeTips:       ageTips,
			Age:           2,
			Status:        v.Status,
			UserType:      v.UserType,
			UsserTypeText: v.UserType.String(),
		})
	}
	return resp
}
