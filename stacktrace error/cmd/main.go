package main

import (
	"github.com/pkg/errors" //this has to be the imported package for errors so that they have a stack
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"io"
	"os"
)

var log zerolog.Logger

type (
	Options struct {
		Level  int8
		Writer io.Writer
	}
)

const (
	// DebugLevel defines debug log level.
	DebugLevel int8 = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
	FatalLevel
	// PanicLevel defines panic log level.
	PanicLevel
	// NoLevel defines an absent log level.
	NoLevel
	// Disabled disables the logger.
	Disabled
	// TraceLevel defines trace log level.
	TraceLevel int8 = -1
)

func init() {
	options := Options{
		Level:  TraceLevel,
		Writer: os.Stdout,
	}
	log = zerolog.New(options.Writer).Level(zerolog.Level(options.Level))
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
}

func main() {
	err := outer()
	log.Error().Stack().Err(err).Msg("")
}

func inner() error {
	return errors.New("seems we have an error here")
}

func middle() error {
	err := inner()
	if err != nil {
		return err
	}
	return nil
}

func outer() error {
	err := middle()
	if err != nil {
		return err
	}
	return nil
}
