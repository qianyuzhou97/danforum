package web

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/database"
)

// type PostService struct {
// 	db *sqlx.DB
// }

func (s *Server) ListAllPosts(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	list, err := s.DB.ListAllPosts(ctx)

	if err != nil {
		return errors.Wrap(err, "error: selecting posts")
	}

	return Respond(ctx, w, list, http.StatusOK)
}

func (s *Server) GetPostByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	list, err := s.DB.GetPostByID(ctx, id)

	if err != nil {
		return errors.Wrap(err, "error: get posts by ID")
	}

	return Respond(ctx, w, list, http.StatusOK)
}

func (s *Server) CreatePost(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var np database.NewPost
	if err := Decode(r, &np); err != nil {
		return errors.Wrap(err, "error decoding post")
	}

	if err := s.DB.CreatePost(ctx, np); err != nil {
		return errors.Wrap(err, "error creating post")
	}
	return nil
}

func (s *Server) UpdatePostByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	var update database.UpdatePost
	if err := Decode(r, &update); err != nil {
		return errors.Wrap(err, "decoding post update")
	}

	if err := s.DB.UpdatePostByID(ctx, update); err != nil {
		return errors.Wrapf(err, "updating post %q", update.ID)
	}

	return nil
}

// Delete removes a single product identified by an ID in the request URL.
func (s *Server) DeletePostByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	if err := s.DB.DeletePostByID(ctx, id); err != nil {
		return errors.Wrapf(err, "deleting post %q", id)
	}

	return Respond(ctx, w, nil, http.StatusNoContent)
}
