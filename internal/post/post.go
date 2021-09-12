package post

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func ListAll(ctx context.Context, db *sqlx.DB) ([]Post, error) {
	post := []Post{}

	const q = `SELECT * FROM posts`

	if err := db.SelectContext(ctx, &post, q); err != nil {
		return nil, errors.Wrap(err, "selecting posts")
	}
	return post, nil
}

func GetByID(ctx context.Context, db *sqlx.DB, postID int64) (*Post, error) {
	var p Post

	const q = `SELECT * FROM posts WHERE post_id = ?`

	if err := db.GetContext(ctx, &p, q, postID); err != nil {
		return nil, errors.Wrap(err, "error get posts based on post_id")
	}
	return &p, nil
}

func CreateNewPost(ctx context.Context, db *sqlx.DB, np NewPost) error {
	const q = `insert into posts(post_id, title, content, author_id, community_id) 
				values(?,?,?,?,?)`

	if _, err := db.ExecContext(ctx, q, np.ID, np.Title, np.Content, np.Author, np.Community); err != nil {
		return errors.Wrap(err, "error get posts based on post_id")
	}
	return nil
}

func UpdateByID(ctx context.Context, db *sqlx.DB, up UpdatePost) error {
	p, err := GetByID(ctx, db, up.ID)
	if err != nil {
		return err
	}

	if up.Title != "" {
		p.Title = up.Title
	}

	if up.Content != "" {
		p.Content = up.Content
	}

	const q = `update posts set title = ?,content = ? where post_id = ?`
	if _, err = db.ExecContext(ctx, q, p.Title, p.Content, p.ID); err != nil {
		return errors.Wrap(err, "updating post")
	}
	return nil
}

func DeleteByID(ctx context.Context, db *sqlx.DB, postID string) error {

	const q = `DELETE FROM posts WHERE post_id = ?`

	if _, err := db.ExecContext(ctx, q, postID); err != nil {
		return errors.Wrapf(err, "deleting post %s", postID)
	}

	return nil
}
