package login

var Svc ILoginSvc

func Init(url string) {
	Svc = NewLoginSvc(url)
}
