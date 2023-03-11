package run

type logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
}
