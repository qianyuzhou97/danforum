package handlers

import (
	"context"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/auth"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"github.com/qianyuzhou97/danforum/internal/user"
	"go.uber.org/zap"
)

type UserService struct {
	db    *sqlx.DB
	sugar *zap.SugaredLogger
}

func (u *UserService) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu user.NewUser

	if err := web.Decode(r, &nu); err != nil {
		return errors.Wrap(err, "error decoding new user")
	}

	if err := user.Create(ctx, u.db, nu); err != nil {
		return errors.Wrap(err, "error creating community")
	}
	return nil
}

func (u *UserService) Token(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(web.KeyValues).(*web.Values)
	if !ok {
		return errors.New("web value missing from context")
	}

	username, pass, ok := r.BasicAuth()
	if !ok {
		err := errors.New("must provide email and password in Basic auth")
		return web.NewRequestError(err, http.StatusUnauthorized)
	}

	err := user.Authenticate(ctx, u.db, username, pass)
	if err != nil {
		switch err {
		case user.ErrAuthenticationFailure:
			return web.NewRequestError(err, http.StatusUnauthorized)
		default:
			return errors.Wrap(err, "authenticating")
		}
	}

	var tkn struct {
		Token string `json:"token"`
	}
	tkn.Token, err = auth.GenToken(username, v.Start, time.Hour)
	if err != nil {
		return errors.Wrap(err, "generating token")
	}

	return web.Respond(ctx, w, tkn, http.StatusOK)
}
