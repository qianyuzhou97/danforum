package schema

import (
	"github.com/jmoiron/sqlx"
)

const seeds = `
INSERT INTO posts (post_id, title, content, author_id, community_id) VALUES
	(1, 'one go question', 'test content', 1, 1),
	(2, 'two go question', 'test content 2', 2, 2)

	INSERT INTO community (community_id, name, introduction) VALUES
	(1, 'Go community', 'Go lover place'),
	(2, 'Python community', 'test content')
	
`

//TODO

func Seed(db *sqlx.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	if _, err := tx.Exec(seeds); err != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return tx.Commit()
}
