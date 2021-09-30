package web

import (
	"context"
	"expvar"
	"net/http"
	"runtime"
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/util/auth"
	"go.uber.org/zap"
)

// m contains the global program counters for the application.
var m = struct {
	gr  *expvar.Int
	req *expvar.Int
	err *expvar.Int
}{
	gr:  expvar.NewInt("goroutines"),
	req: expvar.NewInt("requests"),
	err: expvar.NewInt("errors"),
}

// Middleware is a function designed to run some code before and/or after
// another Handler. It is designed to remove boilerplate or other concerns not
// direct to any given Handler.
type Middleware func(Handler) Handler

// wrapMiddleware creates a new handler by wrapping middleware around a final
// handler. The middlewares' Handlers will be executed by requests in the order
// they are provided.
func wrapMiddleware(mw []Middleware, handler Handler) Handler {

	// Loop backwards through the middleware invoking each one. Replace the
	// handler with the new wrapped handler. Looping backwards ensures that the
	// first middleware of the slice is the first to be executed by requests.
	for i := len(mw) - 1; i >= 0; i-- {
		h := mw[i]
		if h != nil {
			handler = h(handler)
		}
	}

	return handler
}

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate() Middleware {

	// This is the actual middleware function to be executed.
	f := func(after Handler) Handler {

		// Wrap this handler around the next one provided.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// Parse the authorization header. Expected header is of
			// the format `Bearer <token>`.
			parts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				err := errors.New("expected authorization header format: Bearer <token>")
				return NewRequestError(err, http.StatusUnauthorized)
			}

			claims, err := auth.ParseToken(parts[1])
			if err != nil {
				return NewRequestError(err, http.StatusUnauthorized)
			}

			// Add claims to the context so they can be retrieved later.
			ctx = context.WithValue(ctx, auth.Key, claims)

			return after(ctx, w, r)
		}

		return h
	}

	return f
}

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(sugar *zap.SugaredLogger) Middleware {

	// This is the actual middleware function to be executed.
	f := func(before Handler) Handler {

		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			// Run the handler chain and catch any propagated error.
			if err := before(ctx, w, r); err != nil {

				// Log the error.
				sugar.Error(err)

				// Respond to the error.
				if err := RespondError(ctx, w, err); err != nil {
					return err
				}
			}

			// Return nil to indicate the error has been handled.
			return nil
		}

		return h
	}

	return f
}

// Logger writes some information about the request to the logs in the
// format: (200) GET /foo -> IP ADDR (latency)
func Logger(sugar *zap.SugaredLogger) Middleware {

	// This is the actual middleware function to be executed.
	f := func(before Handler) Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			err := before(ctx, w, r)
			v, ok := ctx.Value(KeyValues).(*Values)
			if !ok {
				return errors.New("web value missing from context")
			}
			sugar.Infof("(%d) : %s %s -> %s (%s)",
				v.StatusCode,
				r.Method, r.URL.Path,
				r.RemoteAddr, time.Since(v.Start),
			)

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return f
}

// Metrics updates program counters.
func Metrics() Middleware {

	// This is the actual middleware function to be executed.
	f := func(before Handler) Handler {

		// Wrap this handler around the next one provided.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			err := before(ctx, w, r)

			// Increment the request counter.
			m.req.Add(1)

			// Update the count for the number of active goroutines every 100 requests.
			if m.req.Value()%100 == 0 {
				m.gr.Set(int64(runtime.NumGoroutine()))
			}

			// Increment the errors counter if an error occurred on this request.
			if err != nil {
				m.err.Add(1)
			}

			// Return the error so it can be handled further up the chain.
			return err
		}

		return h
	}

	return f
}
