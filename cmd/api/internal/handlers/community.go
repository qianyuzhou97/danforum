package handlers

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/community"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"go.uber.org/zap"
)

type CommunityService struct {
	db    *sqlx.DB
	sugar *zap.SugaredLogger
}

func (c *CommunityService) ListAll(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	list, err := community.ListAll(ctx, c.db)

	if err != nil {
		return errors.Wrap(err, "error: selecting community")
	}

	return web.Respond(ctx, w, list, http.StatusOK)
}

func (c *CommunityService) GetByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	list, err := community.GetByID(ctx, c.db, id)

	if err != nil {
		return errors.Wrap(err, "error: get community by ID")
	}

	return web.Respond(ctx, w, list, http.StatusOK)
}

func (c *CommunityService) Create(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nc community.NewCommunity
	if err := web.Decode(r, &nc); err != nil {
		return errors.Wrap(err, "error decoding community")
	}

	if err := community.Create(ctx, c.db, nc); err != nil {
		return errors.Wrap(err, "error creating community")
	}
	return nil
}
