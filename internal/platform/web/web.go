package web

import (
	"net/http"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	sugar *zap.SugaredLogger
	mux   *chi.Mux
}

type Handler func(w http.ResponseWriter, r *http.Request) error

func NewApp(sugar *zap.SugaredLogger) *App {
	return &App{
		sugar: sugar,
		mux:   chi.NewRouter(),
	}
}

func (a *App) Handle(method, url string, h Handler) {

	fn := func(w http.ResponseWriter, r *http.Request) {

		// Call the handler and catch any propagated error.
		err := h(w, r)

		if err != nil {

			// Log the error.
			a.sugar.Error(err)

			// Respond to the error.
			if err := RespondError(w, err); err != nil {
				a.sugar.Error(err)
			}
		}
	}

	a.mux.MethodFunc(method, url, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
