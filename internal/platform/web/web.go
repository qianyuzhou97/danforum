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

func NewApp(sugar *zap.SugaredLogger) *App {
	return &App{
		sugar: sugar,
		mux:   chi.NewRouter(),
	}
}

func (a *App) Handle(method, url string, h http.HandlerFunc) {
	a.mux.MethodFunc(method, url, h)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
