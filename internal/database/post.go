package database

import (
	"context"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/util/auth"
	"github.com/qianyuzhou97/danforum/internal/util/snowflake"
)

func (d *DB) ListAllPosts(ctx context.Context) ([]Post, error) {
	post := []Post{}

	const q = `SELECT * FROM posts`

	if err := d.DB.SelectContext(ctx, &post, q); err != nil {
		return nil, errors.Wrap(err, "selecting posts")
	}
	return post, nil
}

func (d *DB) GetPostByID(ctx context.Context, postID int64) (*Post, error) {
	var p Post

	const q = `SELECT * FROM posts WHERE post_id = ?`

	if err := d.DB.GetContext(ctx, &p, q, postID); err != nil {
		return nil, errors.Wrap(err, "error get posts based on post_id")
	}
	return &p, nil
}

func (d *DB) CreatePost(ctx context.Context, np NewPost) error {
	claims, ok := ctx.Value(auth.Key).(auth.Claims)
	if !ok {
		return errors.New("claims missing from context")
	}
	const q = `insert into posts(post_id, title, content, author_id, community_id) 
				values(?,?,?,?,?)`

	if _, err := d.DB.ExecContext(ctx, q, snowflake.GenID(), np.Title, np.Content, claims.Username, 1); err != nil {
		return errors.Wrap(err, "error get posts based on post_id")
	}
	return nil
}

func (d *DB) UpdatePostByID(ctx context.Context, up UpdatePost) error {
	p, err := d.GetPostByID(ctx, up.ID)
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
	if _, err = d.DB.ExecContext(ctx, q, p.Title, p.Content, p.ID); err != nil {
		return errors.Wrap(err, "updating post")
	}
	return nil
}

func (d *DB) DeletePostByID(ctx context.Context, postID string) error {

	const q = `DELETE FROM posts WHERE post_id = ?`

	if _, err := d.DB.ExecContext(ctx, q, postID); err != nil {
		return errors.Wrapf(err, "deleting post %s", postID)
	}

	return nil
}
