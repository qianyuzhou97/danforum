package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/database"
	"github.com/qianyuzhou97/danforum/internal/schema"
	"go.uber.org/zap"
)

// Variables that should be set based on environment variables from Kuberentes
var (
	username = flag.String("username", "root", "username for MySQL")
	password = flag.String("password", "root", "password for MySQL")
	dbName   = flag.String("dbname", "danforum", "database used in MySQL")
)

func main() {
	//initialize zap logger
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer logger.Sync()
	sugar := logger.Sugar()

	if err := run(sugar); err != nil {
		sugar.Fatal("shutting down", "error:", err)
	}
}

func run(sugar *zap.SugaredLogger) error {

	sugar.Info("main : Started")
	defer sugar.Info("main : Completed")

	// MySQL database set up
	db, err := database.Open(*username, *password, *dbName)
	if err != nil {
		return errors.Wrap(err, "connecting to database")
	}
	defer db.Close()

	flag.Parse()

	switch flag.Arg(0) {
	case "migrate":
		if err := schema.Migrate(db); err != nil {
			return errors.Wrap(err, "error applying migrations")
		}
		sugar.Info("Migrations complete")

	case "seed":
		if err := schema.Seed(db); err != nil {
			return errors.Wrap(err, "error seeding database")
		}
		sugar.Info("Seed data complete")
	}
	return nil
}
