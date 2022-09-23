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

var Infof = logger.Info().Msgf

func Debug(v ...interface{}) {
    logger.Debug().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

var Debugf = logger.Debug().Msgf

var Print = Debug
var Printf = Debugf

func Error(v ...interface{}) {
    logger.Error().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

var Errorf = logger.Error().Msgf

func Fatal(v ...interface{}) {
    logger.Fatal().CallerSkipFrame(1).Msg(fmt.Sprint(v...))
}

var Fatalf = logger.Fatal().Msgf
