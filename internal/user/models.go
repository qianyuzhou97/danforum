package user

import "time"

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
