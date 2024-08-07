package mylogin

var (
	Svc ILoginSvc = (*LoginSvc)(nil)
)

func Init(url string) {
	Svc = NewLoginSvc(url)
}
