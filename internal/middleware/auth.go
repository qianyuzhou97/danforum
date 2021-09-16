package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/auth"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
)

// Authenticate validates a JWT from the `Authorization` header.
func Authenticate() web.Middleware {

	// This is the actual middleware function to be executed.
	f := func(after web.Handler) web.Handler {

		// Wrap this handler around the next one provided.
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			// Parse the authorization header. Expected header is of
			// the format `Bearer <token>`.
			parts := strings.Split(r.Header.Get("Authorization"), " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				err := errors.New("expected authorization header format: Bearer <token>")
				return web.NewRequestError(err, http.StatusUnauthorized)
			}

			claims, err := auth.ParseToken(parts[1])
			if err != nil {
				return web.NewRequestError(err, http.StatusUnauthorized)
			}

			// Add claims to the context so they can be retrieved later.
			ctx = context.WithValue(ctx, auth.Key, claims)

			return after(ctx, w, r)
		}

		return h
	}

	return f
}
