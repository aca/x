package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	DebugLevel = zerolog.DebugLevel
	InfoLevel  = zerolog.InfoLevel
	WarnLevel  = zerolog.WarnLevel
	ErrorLevel = zerolog.ErrorLevel
	FatalLevel = zerolog.FatalLevel
	PanicLevel = zerolog.PanicLevel
	NoLevel    = zerolog.NoLevel
	Disabled   = zerolog.Disabled
	TraceLevel = zerolog.TraceLevel
)

func SetLevel(level zerolog.Level) {
	Logger = Logger.Level(level)
}

var Logger = log.Output(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	NoColor:    false,
	TimeFormat: time.RFC3339,
}).With().Caller().Logger()

func Info(v ...interface{}) {
	Logger.Info().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	Logger.Info().CallerSkipFrame(1).Msgf(format, v...)
}

func Debug(v ...interface{}) {
	Logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
	Logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

func Error(v ...interface{}) {
	Logger.Error().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	Logger.Error().CallerSkipFrame(1).Msgf(format, v...)
}

var (
	Print   = Debug
	Println = Debug
	Printf  = Debugf
)

func Fatal(v ...interface{}) {
	Logger.Fatal().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
	Logger.Fatal().CallerSkipFrame(1).Msgf(format, v...)
}
