package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const (
	Name = "token"
)

type Claims struct {
	Username string
	jwt.RegisteredClaims
}

func GenerateJwt(username, secret string) (string, error) {
	claims := Claims{
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedString, err := token.SignedString([]byte(secret))

	return signedString, err
}

func ParseJwt(token, secret string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
