package httpc

type IResponse interface {
	Parse([]byte) error
	Valid() bool
	IsSuccess() bool
	Msg() string
	ParseData() error
}
