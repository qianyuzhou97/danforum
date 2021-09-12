package middleware

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"go.uber.org/zap"
)

// Logger writes some information about the request to the logs in the
// format: (200) GET /foo -> IP ADDR (latency)
func Logger(sugar *zap.SugaredLogger) web.Middleware {

	// This is the actual middleware function to be executed.
	f := func(before web.Handler) web.Handler {

		// Create the handler that will be attached in the middleware chain.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

			err := before(ctx, w, r)
			v, ok := ctx.Value(web.KeyValues).(*web.Values)
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
