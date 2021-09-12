package community

import "time"

type Community struct {
	ID           int64       `db:"community_id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Introduction string    `db:"introduction" json:"introduction"`
	Create_time  time.Time `db:"create_time" json:"create_time"`
	Update_time  time.Time `db:"update_time" json:"update_time"`
}

type NewCommunity struct {
	ID           int64    `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Introduction string `json:"introduction" validate:"required"`
}
