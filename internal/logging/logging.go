package logging

import (
	"os"

	"github.com/rs/zerolog"
)

// New returns a configured zerolog.Logger scoped to the caller.
// It does not mutate zerolog global state.
func New(debug, disable bool) zerolog.Logger {
	level := zerolog.InfoLevel
	if debug {
		level = zerolog.DebugLevel
	}
	if disable {
		level = zerolog.Disabled
	}
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	return zerolog.New(os.Stderr).With().Timestamp().Logger().Level(level)
}
