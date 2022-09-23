package log

import (
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

var Info = logger.Info().Msg
var Infof = logger.Info().Msgf

var Debug = logger.Debug().Msg
var Debugf = logger.Debug().Msgf

var Print = Debug
var Printf = Debugf

var Error = logger.Error().Msg
var Errorf = logger.Error().Msgf

var Fatal = logger.Fatal().Msg
var Fatalf = logger.Fatal().Msgf
