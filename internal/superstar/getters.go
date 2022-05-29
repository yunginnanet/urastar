package superstar

import (
	"io"
	"net/http"

	"git.tcp.direct/kayos/urastar/github"
	"git.tcp.direct/kayos/urastar/internal/config"
)

func (s *SuperStar) MustHandleBody() (stars github.Stars) {
	var result = new(github.Stars)
	bb, rErr := io.ReadAll(s.LatestResp.Body)
	slog := s.hLog(rErr)
	slog.Trace().Str("body", string(bb)).Msg("successful GET body")
	err := result.UnmarshalJSON(bb)
	if err != nil {
		slog.Fatal().Caller(1).Err(err).Msg("failed to read JSON")
	}
	return *result
}

func UserAgent(r *http.Request) *http.Request {
	r.Header.Set("Accept", "application/json")
	r.Header.Set("User-Agent", "yunginnanet/urastar "+config.Version)
	if len(config.GithubToken) > 0 {
		r.Header.Set("Authorization", "token "+config.GithubToken)
	}
	return r
}

func (s *SuperStar) HTTPGet(target string) *http.Response {
	return s.HTTPRequest(http.MethodGet, target)
}

func (s *SuperStar) HTTPHead(target string) *http.Response {
	return s.HTTPRequest(http.MethodHead, target)
}

func (s *SuperStar) HTTPRequest(t, target string) *http.Response {
	var jrq *http.Request
	var err error
	if target == "" {
		target = s.Page
	}
	jrq, err = http.NewRequest(t, target, nil)
	if err != nil {
		s.Log.Panic().Caller().Msg(err.Error())
	}
	if jrq == nil {
		return nil
	}
	jrq = UserAgent(jrq)
	s.LatestResp, err = http.DefaultClient.Do(jrq)
	s.hLog(err).Trace().Msg("Get success")
	return s.LatestResp
}
