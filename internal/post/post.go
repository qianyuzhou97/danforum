package post

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func ListAll(db *sqlx.DB) ([]Post, error) {
	post := []Post{}

	const q = `SELECT * FROM posts`

	if err := db.Select(&post, q); err != nil {
		return nil, errors.Wrap(err, "selecting posts")
	}
	return post, nil
}

func GetByID(db *sqlx.DB, post_id string) (*Post, error) {
	var p Post

	const q = `SELECT * FROM posts WHERE post_id = ?`

	if err := db.Get(&p, q, post_id); err != nil {
		return nil, errors.Wrap(err, "error get posts based on post_id")
	}
	return &p, nil
}
