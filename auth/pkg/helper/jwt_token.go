package helper

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
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
	claims, err := parseClaims(token, secretKey)
	if err != nil {
		return 0, err
	}

	if claims.Type != access {
		return 0, errors.New("it is not access_token")
	}

	return claims.Id, nil
}

func ParseToken(token string, secretKey string) (int32, error) {
	clams, err := parseClaims(token, secretKey)
	return clams.Id, err
}

func generateToken(id int32, secretKey string, expiringDuration time.Duration, type_ string) (string, error) {
	claims := &Claims{
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: time.Now().Add(expiringDuration).Unix(),
		},
		Id:   id,
		Type: type_,
	}

	tokenStr, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).
		SignedString([]byte(secretKey))

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func parseClaims(token string, secretKey string) (*Claims, error) {
	claims := new(Claims)
	jwtToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}

	if !jwtToken.Valid {
		return nil, errors.New("invalid token")
	}

	return claims, err
}

type Claims struct {
	*jwt.StandardClaims
	Id   int32
	Type string
}
