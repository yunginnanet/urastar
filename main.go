package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"git.tcp.direct/kayos/urastar/github"
)

type superStar struct {
	name      string
	stars     github.Stars
	log       *zerolog.Logger
	page, max string
	last      *http.Response
}

func init() {
	if len(os.Args) < 2 {
		println("need target username")
		os.Exit(1)
	}
}

func (s *superStar) hLog(err error) *zerolog.Logger {
	switch {
	case err != nil:
		s.log.Fatal().Caller(1).Msg(err.Error())
	case s.last == nil:
		s.log.Fatal().Caller().Msg("nil response")
	case s.last.StatusCode > 0:
		l := s.log.With().
			Str("caller", s.last.Request.URL.String()).
			Int("status", s.last.StatusCode).
			Int64("size", s.last.ContentLength).
			Logger()
		s.log = &l
	case s.last.StatusCode != 200:
		log.Fatal().Msg("bad status code")
	default:
		//
	}
	return s.log
}

// <https://api.github.com/user/49010538/starred?page=2>; rel="next", <https://api.github.com/user/49010538/starred?page=22>; rel="last"

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
		fmt.Fprintf(os.Stdout, "%v", spl)
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

func (s *superStar) mustHandleBody() (stars github.Stars) {
	var result = new(github.Stars)
	bb, rErr := io.ReadAll(s.last.Body)
	slog := s.hLog(rErr)
	// slog.Trace().Str("body", string(bb)).Msg("successful GET body")
	err := result.UnmarshalJSON(bb)
	if err != nil {
		slog.Fatal().Caller(1).Err(err).Msg("failed to read JSON")
	}
	return *result
}

func (s *superStar) get() *http.Response {
	var jrq *http.Request
	var err error
	jrq, err = http.NewRequest(http.MethodGet, s.page, nil)
	if err != nil {
		log.Panic().Caller().Msg(err.Error())
	}
	if jrq == nil {
		return nil
	}
	jrq.Header.Set("Accept", "application/json")
	jrq.Header.Set("User-Agent", "yunginnanet/urastar 0.1")
	s.last, err = http.DefaultClient.Do(jrq)
	s.hLog(err).Trace().Msg("Get success")
	return s.last
}

func (s *superStar) getAllStars() {
	var (
		moreStars github.Stars
		done      bool
	)
	s.hLog(nil).Trace().Interface("headers", s.last.Header).Msg("successful GET headers")
	s.page = s.last.Request.URL.String()
	s.get()
	moreStars = s.mustHandleBody()
	s.stars = append(s.stars, moreStars...)

	for {
		done = s.nextPage()
		switch {
		case done:
			fallthrough
		default:
			s.get()
			moreStars = s.mustHandleBody()
		}
		s.stars = append(s.stars, moreStars...)
		s.hLog(nil).Info().Int("count", len(s.stars)).Msg("successfully added more stars")
		if done {
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	l := zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stdout
		w.NoColor = false
		w.TimeFormat = time.Stamp
	}))
	super := &superStar{
		name: os.Args[1],
		log:  &l,
	}
	target := github.URLStarsFromName(os.Args[1])
	var err error
	super.last, err = http.DefaultClient.Head(target)
	super.hLog(err)
	super.log.Debug().Msg("stage 1 success")
	super.getAllStars()
	if len(super.stars) < 1 {
		log.Fatal().Caller().Msg("no stars found!")
	}
	/*	for _, s := range super.stars {
			ib, err := json.MarshalIndent(s, "", "\t")
			if err != nil {
				super.log.Fatal().Caller().Err(err).Msg("failed to pretty print")
			}
			println(string(ib))
		}
	*/
	super.log.Info().Int("count", len(super.stars)).Msg("fin")
}
