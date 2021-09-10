package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/cmd/api/internal/handlers"
	"github.com/qianyuzhou97/danforum/internal/platform/database"
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

	api := http.Server{
		Addr:         "localhost:8000",
		Handler:      handlers.API(db, sugar),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	serverErrors := make(chan error, 1)

	// Start the service listening for requests.
	go func() {
		sugar.Infof("main : API listening on %s", api.Addr)
		serverErrors <- api.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	// Blocking main and waiting for shutdown.
	select {
	case err := <-serverErrors:
		return errors.Wrap(err, "error: listening and serving")

	case <-shutdown:
		sugar.Info("main : Start shutdown")

		// Give outstanding requests a deadline for completion.
		const timeout = 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()

		// Asking listener to shutdown and load shed.
		err := api.Shutdown(ctx)
		if err != nil {
			sugar.Infof("main : Graceful shutdown did not complete in %v : %v", timeout, err)
			err = api.Close()
		}

		if err != nil {
			return errors.Wrap(err, "main : could not stop server gracefully")
		}
	}
	return nil
}