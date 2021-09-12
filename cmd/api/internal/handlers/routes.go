package handlers

import (
	"net/http"

	"github.com/jmoiron/sqlx"
	mid "github.com/qianyuzhou97/danforum/internal/middleware"
	"github.com/qianyuzhou97/danforum/internal/platform/web"
	"go.uber.org/zap"
)

func API(db *sqlx.DB, sugar *zap.SugaredLogger) http.Handler {

	// mid.Logger(sugar)
	app := web.NewApp(sugar, mid.Logger(sugar), mid.Errors(sugar), mid.Metrics())

	p := PostService{db: db, sugar: sugar}
	c := CommunityService{db: db, sugar: sugar}
	u := UserService{db: db, sugar: sugar}

	app.Handle(http.MethodGet, "/v1/posts", p.ListAll)
	app.Handle(http.MethodPost, "/v1/posts", p.Create)
	app.Handle(http.MethodGet, "/v1/posts/{id}", p.GetByID)
	app.Handle(http.MethodPut, "/v1/posts/{id}", p.UpdateByID)
	app.Handle(http.MethodDelete, "/v1/posts/{id}", p.DeleteByID)

	app.Handle(http.MethodGet, "/v1/community", c.ListAll)
	app.Handle(http.MethodPost, "/v1/community", c.Create)
	app.Handle(http.MethodGet, "/v1/community/{id}", c.GetByID)

	app.Handle(http.MethodPut, "/v1/user", u.Create)

	return app
}
