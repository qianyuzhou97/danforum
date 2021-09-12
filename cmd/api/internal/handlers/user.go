package handlers

import (
	"context"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
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
