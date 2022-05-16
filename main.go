package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/mailru/easyjson"
	"github.com/rs/zerolog"

	"git.tcp.direct/kayos/urastar/github"
)

var log zerolog.Logger

func init() {
	log = zerolog.New(zerolog.NewConsoleWriter(func(w *zerolog.ConsoleWriter) {
		w.Out = os.Stdout
		w.NoColor = false
		w.TimeFormat = time.Stamp
	}))
	if len(os.Args) < 2 {
		println("need target username")
	}
}

func hLog(log zerolog.Logger, r *http.Response, err error) *zerolog.Logger {
	switch {
	case err != nil:
		log.Fatal().Msg(err.Error())
	case r == nil:
		log.Fatal().Msg("nil response")
	case r.StatusCode > 0:
		log = log.With().
			Int("status", r.StatusCode).
			Int64("size", r.ContentLength).
			Logger()
	case r.StatusCode != 200:
		log.Fatal().Msg("bad status code")
	default:
		//
	}
	return &log
}

// <https://api.github.com/user/49010538/starred?page=2>; rel="next", <https://api.github.com/user/49010538/starred?page=22>; rel="last"

func nextPage(linkHeader string) (next string, ok bool) {
	next = ""
	switch {
	case linkHeader == "":
		ok = false
	case !strings.HasPrefix(linkHeader, "<") || !strings.Contains(linkHeader, ">"):
		ok = false
	case !strings.Contains(linkHeader, "rel=\"next\"") &&
		strings.Contains(linkHeader, "rel=\"prev\""):
		ok = true
	default:
		s := strings.TrimPrefix(linkHeader, "<")
		spl := strings.Split(s, ">;")
		if len(spl) < 1 {
			ok = false
		} else {
			next = spl[0]
			ok = true
		}
	}
	return
}

func mustHandleBody(slog *zerolog.Logger, r *http.Response) (stars []github.Star) {
	var result github.Result
	bb, rErr := io.ReadAll(r.Body)
	if rErr != nil {
		slog.Fatal().Err(rErr).Msg("failed to read body")
	}
	err := easyjson.Unmarshal(bb, &result)
	if err != nil {
		slog.Fatal().Err(rErr).Msg("failed to read JSON")
	}
	return result.Stars
}

func getAllStars(slog *zerolog.Logger, r *http.Response) (userStars []github.Star) {
	var done = false
	for {
		next, ok := nextPage(r.Header.Get("Link"))
		switch {
		case !ok:
			slog.Fatal().Msgf("bad header: %s", r.Header.Get("Link"))
		case next == "" && ok:
			done = true
			fallthrough
		default:
			slog = hLog(log, r, nil)
			userStars = append(userStars, mustHandleBody(slog, r)...)
			slog.Info().Int("count", len(userStars)).Msg("successfully added more stars")
		}
		if done {
			break
		}
	}
	return
}

func main() {
	target := github.URLStarsFromName(os.Args[1])
	r, err := http.DefaultClient.Head(target)
	log = log.With().Str("caller", target).Timestamp().Logger()
	slog := hLog(log, r, err)
	slog.Debug().Msg("success")
	r, err = http.DefaultClient.Get(target)
	slog = hLog(log, r, err)
	slog.Debug().Msg("stage 1 success")
	allStars := getAllStars(slog, r)
	if len(allStars) < 1 {
		log.Fatal().Msg("no stars found!")
	}
	for _, s := range allStars {
		ib, err := json.MarshalIndent(s, "", "\t")
		if err != nil {
			slog.Fatal().Err(err).Msg("failed to pretty print")
		}
		println(string(ib))
	}
}
