package database

import "time"

type Community struct {
	ID           int64     `db:"community_id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Introduction string    `db:"introduction" json:"introduction"`
	Create_time  time.Time `db:"create_time" json:"create_time"`
	Update_time  time.Time `db:"update_time" json:"update_time"`
}

type NewCommunity struct {
	// ID           int64    `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Introduction string `json:"introduction" validate:"required"`
}

type Post struct {
	ID      int64  `db:"post_id" json:"id" `
	Title   string `db:"title" json:"title"`
	Content string `db:"content" json:"content"`
	Author  int64  `db:"author_id" json:"author_id"`
	// Community   int64     `db:"community_id" json:"community_id"`
	Create_time time.Time `db:"create_time" json:"create_time"`
	Update_time time.Time `db:"update_time" json:"update_time"`
}

type NewPost struct {
	// ID      int64  `db:"post_id" json:"id" validate:"required"`
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

type User struct {
	ID          int64     `db:"user_id" json:"id"`
	Username    string    `db:"username" json:"username"`
	Password    string    `db:"password" json:"password"`
	Email       string    `db:"email" json:"email"`
	Create_time time.Time `db:"create_time" json:"create_time"`
	Update_time time.Time `db:"update_time" json:"update_time"`
}

type NewUser struct {
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	PasswordConfirm string `json:"password_confirm" validate:"eqfield=Password"`
	Email           string `json:"email" validate:"required"`
}
