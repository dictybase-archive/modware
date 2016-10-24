package validate

import (
	"fmt"

	"gopkg.in/urfave/cli.v1"
)

func ValidateDbArgs(c *cli.Context) error {
	for _, param := range []string{"user", "password", "database", "host"} {
		if !c.IsSet(param) {
			return cli.NewExitError(fmt.Sprintf("argument %s is not given!!!", param), 2)
		}
	}
	return nil
}
