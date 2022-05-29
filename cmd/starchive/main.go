package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"

	"git.tcp.direct/kayos/urastar/internal/config"
)

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
			config.GithubToken = os.Args[i+1]
			os.Args[i+1] = "_"
		default:
			if a != "-v" && a != "--trace" {
				args = append(args, a)
			}
		}
	}
	return args
}

func main() {

}
