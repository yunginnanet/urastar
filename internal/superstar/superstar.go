package superstar

import (
	"net/http"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"git.tcp.direct/kayos/urastar/github"
)

var ErrNoStarsFound = errors.New("no stars found")

type SuperStar struct {
	Name       string
	Stars      github.Stars
	Log        *zerolog.Logger
	Page, Max  string
	LatestResp *http.Response
}

func (s *SuperStar) hLog(err error) *zerolog.Logger {
	switch {
	case err != nil:
		s.Log.Fatal().Caller(1).Msg(err.Error())
	case s.LatestResp == nil:
		s.Log.Fatal().Caller(1).Msg("nil response")
	case s.LatestResp.StatusCode > 0:
		l := s.Log.With().
			Str("URL", s.LatestResp.Request.URL.String()).Logger()
		if zerolog.GlobalLevel() == zerolog.TraceLevel {
			l = l.With().Int("status", s.LatestResp.StatusCode).
				Interface("headers", s.LatestResp.Header).Logger()
		}
		s.Log = &l
	default:
		//
	}
	if s.LatestResp.StatusCode != 200 {
		s.hLog(nil).Fatal().Msgf("bad status code: %d", s.LatestResp.StatusCode)
	}
	return s.Log
}

func nextPage(hd http.Header) (next, last string, done bool) {
	done = true
	linkHeader := hd.Get("Link")

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
			log.Trace().Msgf("Next page: %s", part)
			next = part
			done = false
		case "prev":
			continue
		case "last":
			last = part
		}
	}
	return
}

func (s *SuperStar) getAllStars() {
	var (
		moreStars github.Stars
		done      bool
	)
	s.hLog(nil).Trace().Interface("headers", s.LatestResp.Header).Msg("successful GET headers")
	s.Page = s.LatestResp.Request.URL.String()
	s.HTTPGet("")
	moreStars = s.MustHandleBody()
	s.Stars = append(s.Stars, moreStars...)

	for {
		s.Page, s.Max, done = nextPage(s.LatestResp.Header)
		switch {
		case done:
			break
		default:
			s.HTTPGet("")
			moreStars = s.MustHandleBody()
		}
		s.Stars = append(s.Stars, moreStars...)
		s.hLog(nil).Debug().Int("count", len(s.Stars)).Msg("successfully retrieved more stars")
		if done {
			break
		}
		time.Sleep(555 * time.Millisecond)
	}
}

func (s *SuperStar) Discover() error {
	var err error
	s.Log.Debug().Msg("getting stars")
	target := github.URLStarsFromName(s.Name)
	s.Log.Trace().Msgf("HEAD %s", target)
	s.HTTPHead(target)
	s.hLog(err)
	s.Log.Trace().Msg("HTTPHead success")
	s.getAllStars()
	if len(s.Stars) < 1 {
		return errors.Wrap(ErrNoStarsFound, s.Name)
	}
	return nil
}
