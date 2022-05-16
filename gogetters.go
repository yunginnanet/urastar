package main

import (
	"io"
	"net/http"

	"git.tcp.direct/kayos/urastar/github"
)

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

func uagent(r *http.Request) *http.Request {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "yunginnanet/urastar 0.1")
	if len(pat) > 0 {
		r.Header.Set("Authorization", "token "+pat)
	}
	return r
}

func (s *superStar) get(target string) *http.Response {
	return s.req(http.MethodGet, target)
}

func (s *superStar) head(target string) *http.Response {
	return s.req(http.MethodHead, target)
}

func (s *superStar) req(t, target string) *http.Response {
	var jrq *http.Request
	var err error
	if target == "" {
		target = s.page
	}
	jrq, err = http.NewRequest(t, target, nil)
	if err != nil {
		s.log.Panic().Caller().Msg(err.Error())
	}
	if jrq == nil {
		return nil
	}
	jrq = uagent(jrq)
	s.last, err = http.DefaultClient.Do(jrq)
	s.hLog(err).Trace().Msg("Get success")
	return s.last
}
