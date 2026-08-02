package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)
func GenerateToken(

	userID int,
	email string,
	secret string,

) (string, error) {

	claims := Claims{

		UserID: userID,
		Email:  email,

		RegisteredClaims: jwt.RegisteredClaims{

			ExpiresAt: jwt.NewNumericDate(
				time.Now().Add(24 * time.Hour),
			),

			IssuedAt: jwt.NewNumericDate(
				time.Now(),
			),

		},
	}

	token := jwt.NewWithClaims(

		jwt.SigningMethodHS256,
		claims,

	)

	return token.SignedString(
		[]byte(secret),
	)
}
func ParseToken(

	tokenString string,
	secret string,

) (*Claims, error) {

	token, err := jwt.ParseWithClaims(

		tokenString,

		&Claims{},

		func(token *jwt.Token) (any, error) {

			return []byte(secret), nil

		},
	)

	if err != nil {

		return nil, err

	}

	claims, ok := token.Claims.(*Claims)

	if !ok || !token.Valid {

		return nil, jwt.ErrTokenInvalidClaims

	}

	return claims, nil

}