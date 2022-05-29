package logger

import (
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"git.tcp.direct/kayos/urastar/internal/config"
)

var once = &sync.Once{}

func GetLogger(superstar string) *zerolog.Logger {
	l := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stderr
		w.NoColor = false
		w.TimeFormat = time.Stamp
	})).With().Timestamp().Str("caller", superstar).Logger()

	once.Do(func() {
		if len(config.GithubToken) > 0 {
			l.Debug().Msg("using API token for authentication")
		}
		l.Trace().Msg("verbose mode enabled")
	})
	return &l
}
