package main

import (
	"flag"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/database"
	"github.com/qianyuzhou97/danforum/internal/schema"
	"go.uber.org/zap"
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

	//database
	db, err := database.Open()
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
