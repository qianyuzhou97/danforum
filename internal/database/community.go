package database

import (
	"context"

	"github.com/pkg/errors"
)

func (d *DB) ListAllCommunity(ctx context.Context) ([]Community, error) {
	community := []Community{}

	const q = `SELECT * FROM community`

	if err := d.DB.SelectContext(ctx, &community, q); err != nil {
		return nil, errors.Wrap(err, "selecting community")
	}
	return community, nil
}

func (d *DB) GetCommunityByID(ctx context.Context, communityID string) (*Community, error) {
	var p Community

	const q = `SELECT * FROM community WHERE community_id = ?`

	if err := d.DB.GetContext(ctx, &p, q, communityID); err != nil {
		return nil, errors.Wrap(err, "error get community based on community_id")
	}
	return &p, nil
}

func (d *DB) CreateCommunity(ctx context.Context, nc NewCommunity) error {
	const q = `insert into community(community_id, name, introduction) 
				values(?,?,?)`

	if _, err := d.DB.ExecContext(ctx, q, nc.ID, nc.Name, nc.Introduction); err != nil {
		return errors.Wrap(err, "error get community based on community_id")
	}
	return nil
}
