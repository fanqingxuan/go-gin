package user

var (
	Svc IUserSvc = (*UserSvc)(nil)
)

func InitUserSvc(url string) {
	Svc = NewUserSvc(url)
}
