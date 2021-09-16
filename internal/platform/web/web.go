package web

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	sugar *zap.SugaredLogger
	mux   *chi.Mux
	mw    []Middleware
}

// ctxKey represents the type of value for the context key.
type ctxKey int

// KeyValues is how request values or stored/retrieved.
const KeyValues ctxKey = 1

// Values carries information about each request.
type Values struct {
	StatusCode int
	Start      time.Time
}

type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

func NewApp(sugar *zap.SugaredLogger, mw ...Middleware) *App {
	return &App{
		sugar: sugar,
		mux:   chi.NewRouter(),
		mw:    mw,
	}
}

func (a *App) Handle(method, url string, h Handler, mw ...Middleware) {

	// First wrap handler specific middleware around this handler.
	h = wrapMiddleware(mw, h)

	// Add the application's general middleware to the handler chain.
	h = wrapMiddleware(a.mw, h)

	fn := func(w http.ResponseWriter, r *http.Request) {

		v := Values{
			Start: time.Now(),
		}
		ctx := context.WithValue(r.Context(), KeyValues, &v)

		// Call the handler and catch any propagated error.
		err := h(ctx, w, r)

		if err != nil {
			// Log the error.
			a.sugar.Error(err)
		}
	}

	a.mux.MethodFunc(method, url, fn)
}

func (a *App) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	a.mux.ServeHTTP(w, r)
}
