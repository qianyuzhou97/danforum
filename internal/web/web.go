package web

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/qianyuzhou97/danforum/internal/database"
	"go.uber.org/zap"
)

type Server struct {
	DB    database.Store
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

func NewServer() *Server {
	return &Server{mux: chi.NewRouter()}
}

func (s *Server) SetRoutes() *Server {
	s.routes()
	return s
}

func (s *Server) SetLogger(sugar *zap.SugaredLogger) *Server {
	s.sugar = sugar
	return s
}

func (s *Server) SetDB(DB database.Store) *Server {
	s.DB = DB
	return s
}

func (s *Server) Handle(method, url string, h Handler, mw ...Middleware) {

	// First wrap handler specific middleware around this handler.
	h = wrapMiddleware(mw, h)

	// Add the application's general middleware to the handler chain.
	h = wrapMiddleware(s.mw, h)

	fn := func(w http.ResponseWriter, r *http.Request) {

		v := Values{
			Start: time.Now(),
		}
		ctx := context.WithValue(r.Context(), KeyValues, &v)

		// Call the handler and catch any propagated error.
		err := h(ctx, w, r)

		if err != nil {
			// Log the error.
			s.sugar.Error(err)
		}
	}

	s.mux.MethodFunc(method, url, fn)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
