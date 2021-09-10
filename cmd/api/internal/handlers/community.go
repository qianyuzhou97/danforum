package handlers

import (
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

func (c *CommunityService) ListAll(w http.ResponseWriter, r *http.Request) error {

	list, err := community.ListAll(r.Context(), c.db)

	if err != nil {
		return errors.Wrap(err, "error: selecting community")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (c *CommunityService) GetCommunityByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	list, err := community.GetByID(r.Context(), c.db, id)

	if err != nil {
		return errors.Wrap(err, "error: get community by ID")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (c *CommunityService) CreateCommunity(w http.ResponseWriter, r *http.Request) error {
	var nc community.NewCommunity
	if err := web.Decode(r, &nc); err != nil {
		return errors.Wrap(err, "error decoding community")
	}

	if err := community.CreateNewCommunity(r.Context(), c.db, nc); err != nil {
		return errors.Wrap(err, "error creating community")
	}
	return nil
}
