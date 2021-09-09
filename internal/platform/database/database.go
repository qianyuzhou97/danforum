package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func Open() (*sqlx.DB, error) {
	return sqlx.Open("mysql", "root:root@/danforum?parseTime=true")
}
