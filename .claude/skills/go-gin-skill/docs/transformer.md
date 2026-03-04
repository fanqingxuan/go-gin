# Transformer 数据转换

`transformer/` 目录负责将 entity 转换为 typing 响应结构，解耦数据层和接口层。

## 示例

```go
// transformer/user.go
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
```

## 在 Logic 中使用

```go
func (l *GetUsersLogic) Handle(ctx context.Context, req typing.ListReq) (*typing.ListResp, error) {
    var users []*entity.User
    err := dao.User.Ctx(ctx).All(&users)
    if err != nil {
        return nil, err
    }
    return &typing.ListResp{
        Data: transformer.ConvertUserToListData(users),
    }, nil
}
```

## 命名规范

- 文件名：与模块对应，如 `user.go`、`order.go`
- 函数名：`Convert{Entity}To{Target}`，如 `ConvertUserToListData`、`ConvertOrderToDetail`
