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

type PostService struct {
	db    *sqlx.DB
	sugar *zap.SugaredLogger
}

func (p *PostService) ListAll(w http.ResponseWriter, r *http.Request) error {

	list, err := post.ListAll(r.Context(), p.db)

	if err != nil {
		return errors.Wrap(err, "error: selecting posts")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (p *PostService) GetByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	list, err := post.GetByID(r.Context(), p.db, id)

	if err != nil {
		return errors.Wrap(err, "error: get posts by ID")
	}

	return web.Respond(w, list, http.StatusOK)
}

func (p *PostService) Create(w http.ResponseWriter, r *http.Request) error {
	var np post.NewPost
	if err := web.Decode(r, &np); err != nil {
		return errors.Wrap(err, "error decoding post")
	}

	if err := post.CreateNewPost(r.Context(), p.db, np); err != nil {
		return errors.Wrap(err, "error creating post")
	}
	return nil
}

func (p *PostService) UpdateByID(w http.ResponseWriter, r *http.Request) error {

	var update post.UpdatePost
	if err := web.Decode(r, &update); err != nil {
		return errors.Wrap(err, "decoding post update")
	}

	if err := post.UpdateByID(r.Context(), p.db, update); err != nil {
		return errors.Wrapf(err, "updating post %q", update.ID)
	}

	return nil
}

// Delete removes a single product identified by an ID in the request URL.
func (p *PostService) DeleteByID(w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	if err := post.DeleteByID(r.Context(), p.db, id); err != nil {
		return errors.Wrapf(err, "deleting post %q", id)
	}

	return web.Respond(w, nil, http.StatusNoContent)
}
