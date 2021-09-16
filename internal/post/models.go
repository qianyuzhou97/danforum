package post

import "time"

type Post struct {
	ID          int64     `db:"post_id" json:"id" `
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	Author      int64     `db:"author_id" json:"author_id"`
	Community   int64     `db:"community_id" json:"community_id"`
	Create_time time.Time `db:"create_time" json:"create_time"`
	Update_time time.Time `db:"update_time" json:"update_time"`
}

type NewPost struct {
	ID      int64  `db:"post_id" json:"id" validate:"required"`
	Title   string `db:"title" json:"title" validate:"required"`
	Content string `db:"content" json:"content" validate:"required"`
	Author  int64  `db:"author_id" json:"author_id" validate:"omitempty"`
	// Community int    `db:"community_id" json:"community_id" validate:"required"`
}

type UpdatePost struct {
	ID      int64  `db:"post_id" json:"id" validate:"required"`
	Title   string `db:"title" json:"title" validate:"omitempty"`
	Content string `db:"content" json:"content" validate:"omitempty"`
}
