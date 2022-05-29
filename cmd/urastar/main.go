package main

import (
	"os"
	"strings"

	"github.com/rs/zerolog"

	"git.tcp.direct/kayos/urastar/internal/config"
	"git.tcp.direct/kayos/urastar/internal/logger"
	starz "git.tcp.direct/kayos/urastar/internal/superstar"
	"git.tcp.direct/kayos/urastar/internal/util"
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
		case a == "-p" || a == "--pretty":
			config.Pretty = true
		default:
			if a != "-v" && a != "--trace" {
				args = append(args, a)
			}
		}
	}
	if len(args) < 1 {
		println("\t\t\t⭐⭐⭐⭐⭐ur a star!!11⭐⭐⭐⭐⭐\n\nUsage: \n\n" + os.Args[0] + "[-t <gh pat>] [--pretty] [-q/-v] <target> [target] [target] [...]\n")
		os.Exit(1)
	}
	return args
}

func main() {
	for _, superstar := range getArgs() {
		super := &starz.SuperStar{
			Name: superstar,
			Log:  logger.GetLogger(superstar),
		}

		err := super.Discover()
		if err != nil {
			super.Log.Error().Msg(err.Error())
			continue
		}

		for _, star := range super.Stars {
			switch config.Pretty {
			case true:
				if err := util.PrettyPrint(star); err != nil {
					println(err.Error())
					return
				}
			default:
				if err := util.JSONPrint(star); err != nil {
					println(err.Error())
					return
				}
			}
		}
		print("\n")

		super.Log.Info().Int("count", len(super.Stars)).Msg("fin")
	}
}
