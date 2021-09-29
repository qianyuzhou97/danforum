package web

import (
	"context"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/database"

)

// type CommunityService struct {
// 	db *sqlx.DB
// }

func (s *Server) ListAllCommunity(ctx context.Context, w http.ResponseWriter, r *http.Request) error {

	list, err := s.DB.ListAllCommunity(ctx)

	if err != nil {
		return errors.Wrap(err, "error: selecting community")
	}

	return Respond(ctx, w, list, http.StatusOK)
}

func (s *Server) GetCommunityByID(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	id := chi.URLParam(r, "id")

	list, err := s.DB.GetCommunityByID(ctx, id)

	if err != nil {
		return errors.Wrap(err, "error: get community by ID")
	}

	return Respond(ctx, w, list, http.StatusOK)
}

func (s *Server) CreateCommunity(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
	var nc database.NewCommunity
	if err := Decode(r, &nc); err != nil {
		return errors.Wrap(err, "error decoding community")
	}

	if err := s.DB.CreateCommunity(ctx, nc); err != nil {
		return errors.Wrap(err, "error creating community")
	}
	return nil
}
