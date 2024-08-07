package httpc

type IBaseResponse interface {
	Parse([]byte) error
	Valid() bool
	IsSuccess() bool
	Msg() string
}

type IResponse interface {
	IBaseResponse
	ParseData() error
}

type IRepsonseNonStardard interface {
	IBaseResponse
	ParseData([]byte) error
}

type BaseSvc struct {
	client *Client
}

func NewBaseSvc(url string) *BaseSvc {
	return &BaseSvc{
		client: NewClient().SetBaseURL(url),
	}
}

func (b *BaseSvc) Client() *Client {
	return b.client
}
