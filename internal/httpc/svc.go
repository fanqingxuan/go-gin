package httpc

type IResponse interface {
	Parse([]byte) error
	Valid() bool
	IsSuccess() bool
	Msg() string
	ParseData() error
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
