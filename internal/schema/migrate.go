package schema

import (
	"github.com/GuiaBolso/darwin"
	"github.com/jmoiron/sqlx"
)

var migrations = []darwin.Migration{
	{
		Version:     1,
		Description: "Add posts Table",
		Script: `
CREATE TABLE posts (
	post_id   INT primary key,
	title        VARCHAR(128),
	content      VARCHAR(8192),
	author_id  INT,
	community_id INT, 
	create_time TIMESTAMP default CURRENT_TIMESTAMP,
	update_time TIMESTAMP default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
);`,
	},
	{
		Version:     2,
		Description: "Add community Table",
		Script: `
CREATE TABLE community (
	community_id   INT primary key,
	name        VARCHAR(128),
	introduction      VARCHAR(256),
	create_time TIMESTAMP default CURRENT_TIMESTAMP,
	update_time TIMESTAMP default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP
);`,
	},
	{
		Version:     3,
		Description: "Add User Table",
		Script: `
CREATE TABLE user (
	user_id   INT    primary key,
	username        VARCHAR(128),
	password        VARCHAR(128),
	introduction    VARCHAR(256),
	email       varchar(64)                         null,
    gender      tinyint   default 0                 not null,
	create_time TIMESTAMP default CURRENT_TIMESTAMP,
	update_time TIMESTAMP default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,

	constraint idx_username
	unique (username)
);`,
	},
}

func Migrate(db *sqlx.DB) error {

	driver := darwin.NewGenericDriver(db.DB, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
