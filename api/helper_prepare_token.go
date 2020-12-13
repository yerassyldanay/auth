package api

import (
	"auth/utils/constants"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"os"
)

// errors
var errorTokenIsInvalid = errors.New("token is invalid")

type JwtClaim interface {
	GenerateToken() (string, error)
	DecodeToken(token string) (int, error)
}

// claim -> token
func OnlyGetTokenFromClaim(claim JwtClaim) (string, error) {
	return claim.GenerateToken()
}

// token -> claim
func OnlyDecodeTokenToClaim(claim JwtClaim) (string, error) {
	return claim.GenerateToken()
}

// for sign up with email
type signUpByEnteringEmailClaim struct {
	Address				string
	jwt.StandardClaims
}

func (sec *signUpByEnteringEmailClaim) GenerateToken() (string, error) {
	// get jwt secret key
	jwt_secret_key := os.Getenv(constants.KEY_JWT_SIGN_UP_EMAIL_SECRET_KEY)
	if jwt_secret_key == "" {
		return "", errors.New("could not get a secret key")
	}

	// get token
	token := jwt.NewWithClaims(jwt.SigningMethodES256, sec)

	// jwt string
	tokenString, err := token.SignedString(jwt_secret_key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (sec *signUpByEnteringEmailClaim) DecodeToken(token string) (int, error) {
	tkn, err := jwt.ParseWithClaims(token, sec, func(token *jwt.Token) (interface{}, error) {
		return os.Getenv(constants.KEY_JWT_SIGN_UP_EMAIL_SECRET_KEY), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return http.StatusUnauthorized, err
		}
		return http.StatusBadRequest, err
	}

	if !tkn.Valid {
		return http.StatusUnauthorized, errorTokenIsInvalid
	}

	return http.StatusOK, nil
}
