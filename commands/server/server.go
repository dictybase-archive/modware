package server

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dictyBase/go-middlewares/middlewares/chain"
	cmw "github.com/dictyBase/go-middlewares/middlewares/cors"
	"github.com/dictyBase/go-middlewares/middlewares/logrus"
	"github.com/dictyBase/go-middlewares/middlewares/router"
	"github.com/dictyBase/modware/resources/publication"
	"github.com/dictyBase/modware/routes"
	"github.com/gocraft/dbr"
	"github.com/gocraft/dbr/dialect"
	"github.com/jackc/pgx"
	"github.com/jackc/pgx/stdlib"
	"github.com/rs/cors"
	"gopkg.in/urfave/cli.v1"
)

func RunServer(c *cli.Context) error {

	// logging middleware
	loggerMw, err := getLoggerMiddleware(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	// middleware chain
	baseChain := chain.NewChain(loggerMw.MiddlewareFn, cmw.CorsAdapter(cors.Default()))

	// database handler
	dbh, err := getPgWrapper(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	// http routes
	r := router.NewRouter()
	routes.AddPublication(&publication.Publication{dbh, c.String("version")}, baseChain, r)
	routes.AddAuthor(&publication.Author{dbh, c.String("version")}, baseChain, r)

	// start the server
	log.Printf("Starting modware api web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), r.Router))
	return nil
}

func getPgWrapper(c *cli.Context) (*dbr.Connection, error) {
	var dbh *dbr.Connection
	h, err := getPgHandler(c)
	if err != nil {
		return dbh, err
	}
	return &dbr.Connection{
		DB:            h,
		Dialect:       dialect.PostgreSQL,
		EventReceiver: &dbr.NullEventReceiver{},
	}, nil
}

func getPgHandler(c *cli.Context) (*sql.DB, error) {
	var db *sql.DB
	config := pgx.ConnPoolConfig{
		ConnConfig: pgx.ConnConfig{
			Host:      c.String("host"),
			User:      c.String("user"),
			Password:  c.String("password"),
			Database:  c.String("database"),
			Port:      uint16(c.Int("port")),
			TLSConfig: nil,
		}}
	pool, err := pgx.NewConnPool(config)
	if err != nil {
		return db, err
	}
	db, err = stdlib.OpenFromConnPool(pool)
	if err != nil {
		return db, err
	}
	return db, nil
}

func getLoggerMiddleware(c *cli.Context) (*logrus.Logger, error) {
	var logger *logrus.Logger
	if c.GlobalIsSet("log") {
		w, err := os.Open(c.GlobalString("log"))
		if err != nil {
			return logger, fmt.Errorf("could not open log file for writing %s\n", err)
		}
		if c.GlobalString("logformat") == "json" {
			logger = logrus.NewJSONFileLogger(w)
		} else {
			logger = logrus.NewFileLogger(w)
		}
	} else {
		if c.GlobalString("logformat") == "json" {
			logger = logrus.NewJSONLogger()
		} else {
			logger = logrus.NewLogger()
		}
	}
	return logger, nil
}
