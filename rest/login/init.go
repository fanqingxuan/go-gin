package login

var (
	Svc ILoginSvc = (*LoginSvc)(nil)
)

func Init(url string) {
	Svc = NewLoginSvc(url)
}
