package community

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

func ListAll(ctx context.Context, db *sqlx.DB) ([]Community, error) {
	community := []Community{}

	const q = `SELECT * FROM community`

	if err := db.SelectContext(ctx, &community, q); err != nil {
		return nil, errors.Wrap(err, "selecting community")
	}
	return community, nil
}

func GetByID(ctx context.Context, db *sqlx.DB, communityID string) (*Community, error) {
	var p Community

	const q = `SELECT * FROM community WHERE community_id = ?`

	if err := db.GetContext(ctx, &p, q, communityID); err != nil {
		return nil, errors.Wrap(err, "error get community based on community_id")
	}
	return &p, nil
}

func CreateNewCommunity(ctx context.Context, db *sqlx.DB, nc NewCommunity) error {
	const q = `insert into community(community_id, name, introduction) 
				values(?,?,?)`

	if _, err := db.ExecContext(ctx, q, nc.ID, nc.Name, nc.Introduction); err != nil {
		return errors.Wrap(err, "error get community based on community_id")
	}
	return nil
}
