package util

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/hokaccha/go-prettyjson"
	"github.com/pkg/errors"

	"git.tcp.direct/kayos/urastar/github"
)

func PrettyPrint(star any) error {
	ib, err := prettyjson.Marshal(star)
	if err != nil {
		return errors.Wrap(err, "failed to pretty print")
	}
	_, _ = fmt.Fprint(os.Stdout, string(ib))
	return nil
}

func JSONPrint(content any) error {
	var ib []byte
	var err error
	if data, ok := content.(github.Stars); ok {
		ib, err = data.MarshalJSON()
		if err != nil {
			return errors.Wrap(err, "failed to JSON print using easyjson")
		}
	} else {
		ib, err = json.Marshal(content)
		if err != nil {
			return errors.Wrap(err, "failed to JSON print")
		}
	}
	_, _ = fmt.Fprint(os.Stdout, string(ib))
	return nil
}
