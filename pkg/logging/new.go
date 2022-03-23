package logging

import (
	"fmt"
	"os"

	"go.uber.org/zap"
)

func MustSetup(options ...Option) {
	o := &Options{
		callerSkip: 2,
	}
	for _, opt := range options {
		opt(o)
	}

	logger, err := makeLogger(o)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
		return
	}
	zap.ReplaceGlobals(logger)
}

func makeLogger(o *Options) (*zap.Logger, error) {
	var makeLogger func(options ...zap.Option) (*zap.Logger, error)

	if o.isDebug {
		makeLogger = zap.NewDevelopment
	} else {
		makeLogger = zap.NewProduction
	}

	return makeLogger(
		zap.AddCallerSkip(o.callerSkip),
	)
}
