package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"go.uber.org/zap"
)

func API(db *sqlx.DB, sugar *zap.SugaredLogger) http.Handler {

	app := web.NewApp(sugar)

	p := Posts{db: db, sugar: sugar}

	app.Handle(http.MethodGet, "/v1/posts", p.ListAll)
	app.Handle(http.MethodGet, "/v1/posts/{id}", p.GetByID)

	return app
}
