package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/jmoiron/sqlx"
	"github.com/qianyuzhou97/danforum/internal/post"
	"go.uber.org/zap"
)

type Posts struct {
	db    *sqlx.DB
	sugar *zap.SugaredLogger
}

func (p *Posts) ListAll(w http.ResponseWriter, r *http.Request) {

	list, err := post.ListAll(p.db)

	if err != nil {
		p.sugar.Errorf("error: selecting posts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		p.sugar.Errorf("error when marshaling json, %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.sugar.Errorf("error when writing json, %s", err)
	}
}

func (p *Posts) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	list, err := post.GetByID(p.db, id)

	if err != nil {
		p.sugar.Errorf("error: selecting posts: %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(list)
	if err != nil {
		p.sugar.Errorf("error when marshaling json, %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(data); err != nil {
		p.sugar.Errorf("error when writing json, %s", err)
	}
}
