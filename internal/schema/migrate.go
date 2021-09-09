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
	post_id   INT,
	title        VARCHAR(128),
	content      VARCHAR(8192),
	author_id  INT,
	community_id INT, 
	create_time TIMESTAMP default CURRENT_TIMESTAMP,
	update_time TIMESTAMP default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,

	PRIMARY KEY (post_id)
);`,
	},
}


func Migrate(db *sqlx.DB) error {

	driver := darwin.NewGenericDriver(db.DB, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations, nil)

	return d.Migrate()
}
