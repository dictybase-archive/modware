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
	_ "github.com/lib/pq"
	"github.com/rs/cors"
	"gopkg.in/urfave/cli.v1"
)

func RunServer(c *cli.Context) error {
	// database handler
	dbh, err := getPgHandler(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}

	// logging middleware
	loggerMw, err := getLoggerMiddleware(c)
	if err != nil {
		return cli.NewExitError(err.Error(), 2)
	}
	// middleware chain
	baseChain := chain.NewChain(loggerMw.LoggerMiddlewareFn, cmw.CorsAdapter(cors.Default()))

	// http routes
	r := router.NewRouter()
	routes.AddPublication(&publication.Publication{dbh}, baseChain, r)
	routes.AddAuthor(&publication.Author{dbh}, baseChain, r)

	// start the server
	log.Printf("Starting modware api web server on port %d\n", c.Int("port"))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", c.Int("port")), r.Router))
	return nil
}

func getPgHandler(c *cli.Context) (*sql.DB, error) {
	connString := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable",
		c.String("user"), c.String("password"),
		c.String("database"), c.String("host"))
	return sql.Open("postgres", connString)
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
