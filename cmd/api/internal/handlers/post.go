package handlers

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"github.com/qianyuzhou97/danforum/internal/post"
	"go.uber.org/zap"
)

type Posts struct {
	db    *sqlx.DB
	sugar *zap.SugaredLogger
}

func (p *Posts) ListAll(w http.ResponseWriter, r *http.Request) error {

	list, err := post.ListAll(r.Context(), p.db)

	if err != nil {
		return errors.Wrap(err, "error: selecting posts")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (p *Posts) GetByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	list, err := post.GetByID(r.Context(), p.db, id)

	if err != nil {
		return errors.Wrap(err, "error: get posts by ID")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (p *Posts) CreatePost(w http.ResponseWriter, r *http.Request) error {
	var np post.NewPost
	if err := web.Decode(r, &np); err != nil {
		return errors.Wrap(err, "error decoding post")
	}

	if err := post.CreateNewPost(r.Context(), p.db, np); err != nil {
		return errors.Wrap(err, "error creating post")
	}
	return nil
}
