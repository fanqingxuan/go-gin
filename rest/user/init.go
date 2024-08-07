package user

var (
	Svc IUserSvc = (*UserSvc)(nil)
)

func Init(url string) {
	Svc = NewUserSvc(url)
}
