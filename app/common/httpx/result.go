package httpx

type Result struct {
	Code      int
	Message   string
	Data      any
	RequestId string
}
