package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"go.uber.org/zap"
)

func API(db *sqlx.DB, sugar *zap.SugaredLogger) http.Handler {

	app := web.NewApp(sugar)

	p := PostService{db: db, sugar: sugar}
	c := CommunityService{db: db, sugar: sugar}

	app.Handle(http.MethodGet, "/v1/posts", p.ListAll)
	app.Handle(http.MethodPost, "/v1/posts", p.CreatePost)
	app.Handle(http.MethodGet, "/v1/posts/{id}", p.GetPostByID)

	app.Handle(http.MethodGet, "/v1/community", c.ListAll)
	app.Handle(http.MethodPost, "/v1/community", c.CreateCommunity)
	app.Handle(http.MethodGet, "/v1/community/{id}", c.GetCommunityByID)

	return app
}
