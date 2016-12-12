package main

import (
	"os"

	"github.com/dictyBase/modware/commands/server"
	"github.com/dictyBase/modware/commands/validate"
	"gopkg.in/urfave/cli.v1"
)

func main() {
	app := cli.NewApp()
	app.Name = "modware"
	app.Usage = "A HTTP api server for chado database"
	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "log,l",
			Usage: "Name of the log file(optional), default goes to stderr",
		},
		cli.StringFlag{
			Name:  "logformat,lfmt",
			Usage: "Format of the log(optional), default is json",
			Value: "json",
		},
		cli.StringFlag{
			Name:   "user, u",
			Usage:  "chado database user[REQUIRED]",
			EnvVar: "CHADO_USER",
		},
		cli.StringFlag{
			Name:   "password, p",
			Usage:  "chado database password[REQUIRED]",
			EnvVar: "CHADO_PASS",
		},
		cli.StringFlag{
			Name:   "database, db",
			Usage:  "chado database name[REQUIRED]",
			EnvVar: "CHADO_DB",
		},
		cli.StringFlag{
			Name:   "host, h",
			Usage:  "chado database host[REQUIRED]",
			EnvVar: "CHADO_HOST",
		},
		cli.IntFlag{
			Name:  "port, p",
			Usage: "server port",
			Value: 9998,
		},
	}
	app.Before = validate.ValidateDbArgs
	app.Commands = []cli.Command{
		{
			Name:   "run",
			Usage:  "runs the api server",
			Action: server.RunServer,
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "version, ver",
					Usage: "api version",
					Value: "1.0",
				},
			},
		},
	}
	app.Run(os.Args)
}
