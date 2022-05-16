package main

import (
	"encoding/json"
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
	page, max int
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
		log.Fatal().Msg(err.Error())
	case s.last == nil:
		log.Fatal().Msg("nil response")
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

func nextPage(linkHeader string) (next string) {
	next = ""
	var last = ""
	spl := strings.Split(linkHeader, ",")
	for i, part := range spl {
		part = strings.Trim(part, "<")
		part = strings.Trim(part, "<")
		part = strings.Trim(part, ";")
		part = strings.TrimPrefix(part, "rel=")
		part = strings.Trim(part, "\"")
		part = strings.TrimSpace(part)
		if len(spl) == i+1 {
			break
		}
		switch spl[i+1] {
		case "next":
			next = part
		case "prev":
			continue
		case "last":
			last = part
		}
	}
	if next == last {
		next = ""
	}
	return
}

func (s *superStar) mustHandleBody(r *http.Response) (stars github.Stars, next string) {
	var result = new(github.Stars)
	bb, rErr := io.ReadAll(r.Body)
	slog := s.hLog(rErr)
	// slog.Trace().Str("body", string(bb)).Msg("successful GET body")
	err := result.UnmarshalJSON(bb)
	if err != nil {
		slog.Fatal().Err(err).Msg("failed to read JSON")
	}
	next = nextPage(r.Header.Get("Link"))
	return *result, next
}

func (s *superStar) get(r any) *http.Response {
	var jrq *http.Request
	rs, rok := r.(*http.Response)
	rstr, strok := r.(string)
	switch {
	case rok:
		jrq, _ = http.NewRequest(http.MethodGet, rs.Request.URL.String(), nil)
	case strok:
		jrq, _ = http.NewRequest(http.MethodGet, rstr, nil)
	default:
		log.Panic().Interface("caller", r).Msg("bad type")
	}
	if jrq == nil {
		return nil
	}
	jrq.Header.Set("Accept", "application/json")
	jrq.Header.Set("User-Agent", "yunginnanet/urastar 0.1")
	var err error
	s.last, err = http.DefaultClient.Do(jrq)
	s.hLog(err).Trace().Msg("Get success")
	return s.last
}

func (s *superStar) getAllStars(r *http.Response) (userStars []github.Stars) {
	var (
		done      = false
		moreStars github.Stars
		next      string
		ok        bool
	)
	s.hLog(nil).Trace().Interface("headers", r.Header).Msg("successful GET headers")
	moreStars, next = s.mustHandleBody(s.get(r))
	userStars = append(userStars, moreStars)
	if !ok || next == "" {
		return userStars
	}
	for {
		switch {
		case next == "":
			done = true
		default:
			if next != "" {
				moreStars, next = s.mustHandleBody(s.get(next))
			}
			userStars = append(userStars, moreStars)
			s.hLog(nil).Info().Int("count", len(userStars)).Msg("successfully added more stars")
		}
		if done {
			break
		}
	}
	return
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
	r, err := http.DefaultClient.Head(target)
	super.hLog(err)
	super.log.Debug().Msg("stage 1 success")
	allStars := super.getAllStars(super.get(r))
	if len(allStars) < 1 {
		log.Fatal().Caller().Msg("no stars found!")
	}
	for _, s := range allStars {
		ib, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			super.log.Fatal().Caller().Err(err).Msg("failed to pretty print")
		}
		println(string(ib))
	}
}
