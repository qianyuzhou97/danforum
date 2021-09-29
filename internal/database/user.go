package database

import (
	"context"
	"crypto/md5"
	"database/sql"
	"encoding/hex"

	"github.com/pkg/errors"
	"github.com/qianyuzhou97/danforum/internal/platform/snowflake"
)

const secret = "dan"

var (
	// ErrAuthenticationFailure occurs when a user attempts to authenticate but
	// anything goes wrong.
	ErrAuthenticationFailure = errors.New("Authentication failed")
)

func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

// Create inserts a new user into the database.
func (d *DB) CreateUser(ctx context.Context, n NewUser) error {

	const q = `INSERT INTO user
		(user_id, username, password, email)
		VALUES (?,?,?,?)`
	_, err := d.DB.ExecContext(
		ctx, q, snowflake.GenID(), n.Username, encryptPassword(n.Password), n.Email)
	if err != nil {
		return errors.Wrap(err, "inserting user")
	}

	return nil
}

func (d *DB) Authenticate(ctx context.Context, username, password string) error {

	const q = `SELECT * FROM user WHERE username = ?`

	var u User
	if err := d.DB.GetContext(ctx, &u, q, username); err != nil {

		// Normally we would return ErrNotFound in this scenario but we do not want
		// to leak to an unauthenticated user which emails are in the system.
		if err == sql.ErrNoRows {
			return ErrAuthenticationFailure
		}

		return errors.Wrap(err, "selecting single user")
	}

	// Compare the provided password with the saved hash. Use the bcrypt
	// comparison function so it is cryptographically secure.
	encrypted := encryptPassword(password)
	if encrypted != u.Password {
		return ErrAuthenticationFailure
	}

	return nil
}
