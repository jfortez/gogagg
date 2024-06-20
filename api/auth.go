package api

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Auth struct {
	expiration time.Duration
	secretKey  []byte
}

func NewAuth(expiration time.Duration, secretKey []byte) *Auth {
	return &Auth{
		expiration: expiration,
		secretKey:  secretKey,
	}
}

func (a *Auth) CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": username,
			"exp":  time.Now().Add(a.expiration).Unix(),
		})

	tokenString, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (a *Auth) parseFunc(token *jwt.Token) (interface{}, error) {
	return a.secretKey, nil
}

func (a *Auth) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, a.parseFunc)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
