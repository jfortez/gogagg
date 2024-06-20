package api

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jfortez/gogagg/model"
)

type Session struct {
	expiration time.Duration
	secretKey  []byte
}

func NewSession(expiration time.Duration, secretKey []byte) *Session {
	return &Session{
		expiration: expiration,
		secretKey:  secretKey,
	}
}

func (a *Session) CreateToken(user string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"user": user,
			"exp":  time.Now().Add(a.expiration).Unix(),
		})

	tokenString, err := token.SignedString(a.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (a *Session) parseFunc(token *jwt.Token) (interface{}, error) {
	return a.secretKey, nil
}

func (a *Session) VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, a.parseFunc)

	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}

func (a *Session) getUser(r *http.Request) (model.AuthUser, error) {

	tokenCookie, err := r.Cookie("token")

	if err != nil {
		return model.AuthUser{}, err
	}

	token, err := jwt.Parse(tokenCookie.Value, a.parseFunc)

	if err != nil {
		return model.AuthUser{}, err
	}

	if !token.Valid {
		return model.AuthUser{}, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return model.AuthUser{}, errors.New("invalid token")
	}

	user, ok := claims["user"].(string)

	if !ok {
		return model.AuthUser{}, errors.New("invalid token")
	}

	var userModel model.AuthUser
	err = json.Unmarshal([]byte(user), &userModel)
	if err != nil {
		return model.AuthUser{}, err
	}

	return userModel, nil
}

func (a *Session) removeSession(w http.ResponseWriter, r *http.Request) {
	cToken := &http.Cookie{
		Name:    "token",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	cUser := &http.Cookie{
		Name:    "user",
		Value:   "",
		Path:    "/",
		Expires: time.Unix(0, 0),
	}
	http.SetCookie(w, cUser)
	http.SetCookie(w, cToken)
}

func (a *Session) setSession(w http.ResponseWriter, r *http.Request, user model.AuthUser) error {

	jsonUser, err := json.Marshal(user)
	if err != nil {
		return err
	}

	token, err := a.CreateToken(string(jsonUser))

	if err != nil {
		return err
	}

	cookie := &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Expires:  time.Now().Add(a.expiration),
		Secure:   true,
		HttpOnly: true,
	}
	usrCookie := &http.Cookie{
		Name:     "user",
		Value:    string(jsonUser),
		Path:     "/",
		Expires:  time.Now().Add(a.expiration),
		Secure:   true,
		HttpOnly: true,
	}

	http.SetCookie(w, cookie)
	http.SetCookie(w, usrCookie)

	return nil
}
