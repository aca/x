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

func Infof(format string, v ...interface{}) {
	logger.Info().Msgf(format, v...)
	fmt.Print()
}

func Info(v ...interface{}) {
	logger.Info().Msg(fmt.Sprint(v...))
}

func Debug(v ...interface{}) {
	logger.Print(v...)
}

var Print = Debug

func Debugf(format string, v ...interface{}) {
	logger.Printf(format, v...)
}

var Printf = Debugf

func Error(v ...interface{}) {
	logger.Error().Msg(fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	logger.Error().Msgf(format, v...)
}

func Fatalf(format string, v ...interface{}) {
	logger.Fatal().Msgf(format, v...)
}

func Fatal(v ...interface{}) {
	logger.Fatal().Msg(fmt.Sprint(v...))
}
