package user

var Svc IUserSvc

func Init(url string) {
	Svc = NewUserSvc(url)
}
