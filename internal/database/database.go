package database

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	DB *sqlx.DB
}

func Open(user string, password string, dbName string) (*sqlx.DB, error) {
	//return sqlx.Open("mysql", user+":"+password+"@(mysql:3306)/"+dbName+"?parseTime=true")
	return sqlx.Open("mysql", "root:root@/danforum?parseTime=true")
}
