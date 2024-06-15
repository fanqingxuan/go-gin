package logx

type Logger interface {
	Debug(keyword string, message any)
	Debugf(keyword string, format string, message ...any)

	Info(keyword string, message any)
	Infof(keyword string, format string, message ...any)

	Warn(keyword string, message any)
	Warnf(keyword string, format string, message ...any)

	Error(keyword string, message any)
	Errorf(keyword string, format string, message ...any)
}
