package post

import "time"

type Post struct {
	ID          int       `db:"post_id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Content     string    `db:"content" json:"content"`
	Author      int       `db:"author_id" json:"author"`
	Community   int       `db:"community_id" json:"community"`
	Create_time time.Time `db:"create_time" json:"create_time"`
	Update_time time.Time `db:"update_time" json:"update_time"`
}
