package cli

import (
	"errors"
	"fmt"

	"github.com/urfave/cli/v2"
)

// deploy to azure container apps
// using yaml file as configuration source. yaml is specified with --config flag
func deployAction(cCtx *cli.Context) error {
	if cCtx.IsSet("config") {
		fmt.Println(cCtx.String("config"))
		return nil
	} else {
		return errors.New("config file is not specified")
	}
}
