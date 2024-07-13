package event

type Event interface {
	Name() string
	Payload() interface{}
}
