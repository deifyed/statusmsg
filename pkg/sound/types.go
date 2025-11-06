package sound

type logger interface {
	Warnf(format string, args ...any)
}
