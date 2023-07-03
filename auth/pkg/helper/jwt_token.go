package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"time"
)

const (
	access  = "access"
	refresh = "refresh"
)

type Token struct {
	AccessToken  string
	RefreshToken string
}

func GenerateToken(id int32, secretKey string, accessExpiring, refreshExpiring time.Duration) (*Token, error) {
	token := new(Token)
	var err error

	token.AccessToken, err = generateToken(id, secretKey, accessExpiring, access)
	if err != nil {
		return nil, err
	}

	token.RefreshToken, err = generateToken(id, secretKey, refreshExpiring, refresh)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func ParseAccessToken(token string, secretKey string) (int32, error) {
	claims, id, err := parseClaims(token, secretKey)
	if err != nil {
		return 0, err
	}

	if claims.Subject != access {
		return 0, errors.New("it is not access_token")
	}

	return id, nil
}

func ParseToken(token string, secretKey string) (int32, error) {
	_, id, err := parseClaims(token, secretKey)
	return id, err
}

func generateToken(id int32, secretKey string, expiringDuration time.Duration, subject string) (string, error) {
	claims := &jwt.StandardClaims{
		Id:        strconv.Itoa(int(id)),
		ExpiresAt: time.Now().Add(expiringDuration).Unix(),
		Subject:   subject,
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func parseClaims(token string, secretKey string) (*jwt.StandardClaims, int32, error) {
	claims := new(jwt.StandardClaims)
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, 0, err
	}

	if !jwtToken.Valid {
		return nil, 0, errors.New("invalid token")
	}

	id, err := strconv.Atoi(claims.Id)
	if err != nil {
		return nil, 0, errors.New("modified token")
	}

	return claims, int32(id), err
}
