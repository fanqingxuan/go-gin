package userc

import "go-gin/internal/httpc"

// userc.UserSvc.Hello(ctx,req)

var UserSvc = (*userSvc)(nil)

func InitUserSvc(url string) {
	UserSvc = NewUserSvc(httpc.NewClient().SetBaseURL(url))
}
