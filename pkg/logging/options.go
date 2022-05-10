package logging

type Options struct {
	isDebug    bool
	callerSkip int
}

type Option func(*Options)

func WithCallerSkip(n int) Option {
	return func(o *Options) {
		o.callerSkip = n
	}
}

func WithIsDebug(isDebug bool) Option {
	return func(o *Options) {
		o.isDebug = isDebug
	}
}
