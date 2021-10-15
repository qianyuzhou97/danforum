package auth

import (
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// ctxKey represents the type of value for the context key.
type ctxKey int

// Key is used to store/retrieve a Claims value from a context.Context.
const Key ctxKey = 1

var mySecret = []byte("danforum")

type Claims struct {
	// UserID   int64  `json:"user_id"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// GenToken will generate JWT
func GenToken(username string, now time.Time, expires time.Duration) (string, error) {
	// make a claim
	c := Claims{
		// UserID:   userID,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  now.Unix(),
			ExpiresAt: now.Add(expires).Unix(),
		},
	}
	// specify the encrypt method
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(mySecret)
}

// ParseToken will parse and verify JWT Token
func ParseToken(tokenString string) (*Claims, error) {
	var mc = new(Claims)
	token, err := jwt.ParseWithClaims(tokenString, mc, func(token *jwt.Token) (i interface{}, err error) {
		return mySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if token.Valid {
		return mc, nil
	}
	return nil, errors.New("invalid token")
}
