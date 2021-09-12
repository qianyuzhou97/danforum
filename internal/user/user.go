package user

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/snowflake"
)

const secret = "dan"

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Create inserts a new user into the database.
func Create(ctx context.Context, db *sqlx.DB, n NewUser) error {

	const q = `INSERT INTO user
		(user_id, username, password, email)
		VALUES (?,?,?,?)`
	_, err := db.ExecContext(
		ctx, q, snowflake.GenID(), n.Username, encryptPassword(n.Password), n.Email)
	if err != nil {
		return errors.Wrap(err, "inserting user")
	}

	return nil
}
