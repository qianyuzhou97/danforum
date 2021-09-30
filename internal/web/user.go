package web

import (
	"context"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/database"
	"github.com/qianyuzhou97/danforum/internal/util/auth"
)

// type UserService struct {
// 	db *sqlx.DB
// }

func (s *Server) CreateUser(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nu database.NewUser

	if err := Decode(r, &nu); err != nil {
		return errors.Wrap(err, "error decoding new user")
	}

	if err := s.DB.CreateUser(ctx, nu); err != nil {
		return errors.Wrap(err, "error creating community")
	}
	return nil
}

func (s *Server) Authenticate(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	v, ok := ctx.Value(KeyValues).(*Values)
	if !ok {
		return errors.New("web value missing from context")
	}

	username, pass, ok := r.BasicAuth()
	if !ok {
		err := errors.New("must provide email and password in Basic auth")
		return NewRequestError(err, http.StatusUnauthorized)
	}

	err := s.DB.Authenticate(ctx, username, pass)
	if err != nil {
		switch err {
		case database.ErrAuthenticationFailure:
			return NewRequestError(err, http.StatusUnauthorized)
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

	return Respond(ctx, w, tkn, http.StatusOK)
}
