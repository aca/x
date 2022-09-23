package log

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var logger = log.Output(zerolog.ConsoleWriter{
	Out:        os.Stderr,
	NoColor:    false,
	TimeFormat: time.RFC3339,
}).With().Caller().Logger()

func Info(v ...interface{}) {
    logger.Info().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
    logger.Info().CallerSkipFrame(1).Msgf(format, v...)
}

func Debug(v ...interface{}) {
    logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Debugf(format string, v ...interface{}) {
    logger.Debug().CallerSkipFrame(1).Msgf(format, v...)
}

func Error(v ...interface{}) {
    logger.Error().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
    logger.Error().CallerSkipFrame(1).Msgf(format, v...)
}

var Print = Debug
var Println = Debug
var Printf = Debugf

func Fatal(v ...interface{}) {
    logger.Fatal().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

func Fatalf(format string, v ...interface{}) {
    logger.Fatal().CallerSkipFrame(1).Msgf(format, v...)
}
