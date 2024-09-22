package jwttoken

import (
	"fmt"
	"school-management-app/pkg/config"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

const (
	JwtExpTimeInMinutes = time.Minute * time.Duration(1)
)

var (
	ErrTokenExpired        = fmt.Errorf("jwt token expired")
	ErrTokenIsInvalid      = fmt.Errorf("jwt token is invalid")
	ErrNoCustomClaimsFound = fmt.Errorf("error while parsing token, no custom claims found")
)

type tokenData struct {
	Id       int
	Role     string
	AddlInfo interface{}
}

type CustomClaims struct {
	tokenData
	jwt.RegisteredClaims
}

func GenerateToken(id int, role string, addlInfo interface{}) (string, error) {

	//create a custom claim
	claims := &CustomClaims{
		tokenData: tokenData{
			Id:       id,
			Role:     role,
			AddlInfo: addlInfo,
		},
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(JwtExpTimeInMinutes)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(config.Env.JwtSecretKey))

	return tokenString, err
}

func GetDataFromToken(tokenString string) (tokenData *tokenData, err error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.Env.JwtSecretKey), nil
	})
	if err != nil {
		return nil, fmt.Errorf("jwt token is invalid or error in parsing. (error: %v)", err)
	}
	if !token.Valid {
		return nil, ErrTokenIsInvalid
	}

	if claims, ok := token.Claims.(*CustomClaims); ok {
		//check if token expired
		if claims.ExpiresAt.Time.Before(time.Now()) {
			return nil, ErrTokenExpired
		}
		return &claims.tokenData, nil
	} else {
		return nil, ErrNoCustomClaimsFound
	}
}
