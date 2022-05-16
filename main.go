package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"git.tcp.direct/kayos/urastar/github"
)

var pat string

type superStar struct {
	name      string
	stars     github.Stars
	log       *zerolog.Logger
	page, max string
	last      *http.Response
}

func getArgs() []string {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	var args []string
	for i, a := range os.Args {
		if i == 0 || a == "_" {
			continue
		}
		switch {
		case strings.HasPrefix(a, "-v"), a == "--trace":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		case a == "-q", a == "--quiet":
			zerolog.SetGlobalLevel(zerolog.Disabled)
		case a == "--token" || a == "-t":
			if len(os.Args) == i-1 {
				println("bad syntax, missing token")
				os.Exit(1)
			}
			pat = os.Args[i+1]
			os.Args[i+1] = "_"
		default:
			if a != "-v" && a != "--trace" {
				args = append(args, a)
			}
		}
	}
	if len(args) < 1 {
		println("\t\t\tur a star!!11\n\nUsage: \n\n" + os.Args[0] + "[-t <gh pat>] [-q/-v] <target> [target] [target] [...]\n")
		os.Exit(1)
	}
	return args
}

func (s *superStar) hLog(err error) *zerolog.Logger {
	switch {
	case err != nil:
		s.log.Fatal().Caller(1).Msg(err.Error())
	case s.last == nil:
		s.log.Fatal().Caller(1).Msg("nil response")
	case s.last.StatusCode > 0:
		l := s.log.With().
			Str("caller", s.last.Request.URL.String()).Logger()
		if zerolog.GlobalLevel() == zerolog.TraceLevel {
			l = l.With().Int("status", s.last.StatusCode).
				Interface("headers", s.last.Header).Logger()
		}
		s.log = &l
	default:
		//
	}
	if s.last.StatusCode != 200 {
		s.hLog(nil).Fatal().Msgf("bad status code: %d", s.last.StatusCode)
	}
	return s.log
}

func (s *superStar) nextPage() (ok bool) {
	ok = true
	linkHeader := s.last.Header.Get("Link")

	remove := []string{",", ";", "<", ">", "rel=", "\""}
	for _, rem := range remove {
		linkHeader = strings.ReplaceAll(linkHeader, rem, "")
	}
	spl := strings.Split(linkHeader, " ")
	for i, part := range spl {
		spl[i] = strings.TrimSpace(spl[i])
		if len(spl) == i+1 {
			break
		}
		switch spl[i+1] {
		case "next":
			s.hLog(nil).Trace().Msgf("Next page: %s", part)
			s.page = part
			ok = false
		case "prev":
			continue
		case "last":
			if len(s.max) < 1 {
				s.hLog(nil).Trace().Msgf("Last page: %s", part)
				s.max = part
			}
		}
	}
	return
}

func (s *superStar) getAllStars() {
	var (
		moreStars github.Stars
		done      bool
	)
	s.hLog(nil).Trace().Interface("headers", s.last.Header).Msg("successful GET headers")
	s.page = s.last.Request.URL.String()
	s.get("")
	moreStars = s.mustHandleBody()
	s.stars = append(s.stars, moreStars...)

	for {
		done = s.nextPage()
		switch {
		case done:
			fallthrough
		default:
			s.get("")
			moreStars = s.mustHandleBody()
		}
		s.stars = append(s.stars, moreStars...)
		s.hLog(nil).Debug().Int("count", len(s.stars)).Msg("successfully retrieved more stars")
		if done {
			break
		}
		time.Sleep(555 * time.Millisecond)
	}
}

func (s *superStar) shoot() {
	var err error
	target := github.URLStarsFromName(s.name)
	s.log.Trace().Msgf("HEAD %s", target)
	s.head(target)
	s.hLog(err)
	s.log.Trace().Msg("head success")
	s.getAllStars()
	if len(s.stars) < 1 {
		log.Fatal().Caller().Msg("no stars found!")
	}
	for _, star := range s.stars {
		ib, err := json.MarshalIndent(star, "", "\t")
		if err != nil {
			s.log.Fatal().Caller().Err(err).Msg("failed to pretty print")
		}
		_, _ = fmt.Fprint(os.Stdout, string(ib))
	}
}

func main() {
	once := &sync.Once{}

	for _, superstar := range getArgs() {
		l := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
			w.Out = os.Stderr
			w.NoColor = false
			w.TimeFormat = time.Stamp
		})).With().Timestamp().Logger()
		if len(pat) > 0 {
			log.Debug().Msg("using API token for authentication")
		}
		super := &superStar{
			name: superstar,
			log:  &l,
		}
		once.Do(func() { super.log.Trace().Msg("verbose mode enabled") })
		super.log.Debug().Str("caller", superstar).Msg("getting stars")
		super.shoot()
		super.log.Info().Int("count", len(super.stars)).Msg("fin")
	}
}
