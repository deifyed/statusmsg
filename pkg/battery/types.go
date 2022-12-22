package battery

type logger interface {
	Warnf(format string, args ...interface{})
}
