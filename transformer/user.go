package transformer

import (
	"go-gin/model"
	"go-gin/types"
)

func ConvertUserToListResp(u []model.User) []types.ListResp {
	var resp []types.ListResp
	for _, v := range u {
		var ageTips string
		if *v.Age >= 18 {
			ageTips = "成年"
		} else {
			ageTips = "未成年"
		}
		resp = append(resp, types.ListResp{
			Name:    v.Name,
			Age:     *v.Age,
			AgeTips: ageTips,
		})
	}
	return resp
}
