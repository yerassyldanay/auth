package api

import (
	"auth/utils/helper"
	"auth/utils/randomer"
	"bytes"
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestEmailCreateOne(t *testing.T) {
	emailToSignUp := signUpByEnteringEmail {
		Address:          randomer.RandomOnlyLowerCaseString(20) + "@gmail.com",
	}

	// sign up
	requestBody, err := json.Marshal(&emailToSignUp)
	require.NoError(t, err)

	// new recorder
	recorder := httptest.NewRecorder()
	
	// new request
	request, err := http.NewRequest(http.MethodPost, "/v1/sign-up/email", bytes.NewBuffer(requestBody))
	require.NoError(t, err)

	// send request
	testServer.router.ServeHTTP(recorder, request)
	//fmt.Println("recorder.Code:", recorder.Code)
	//require.Less(t, recorder.Code, 200)

	testParseResponseBody(t, recorder)
}

func TestSignUpConfirmHashFromEmail(t *testing.T) {

	// address
	address := randomer.RandomOnlyLowerCaseString(20) + "@gmail.com"

	// create claim
	claim := signUpByEnteringEmailClaim{
		Address:        address,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: helper.OnlyGetCurrentTime().Add(time.Hour * 1000).Unix(),
			IssuedAt:  helper.OnlyGetCurrentTime().Unix(),
			Issuer: "auth",
		},
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	// Create the JWT string
	tokenString, err := token.SignedString([]byte(testConfiguration.JwtSignUpEmailSecretKey))
	require.NoError(t, err)

	// signUpConfirmHashFromEmail
	su := signUpConfirmHashFromEmail{
		Name:        randomer.RandomOnlyLowerCaseString(20),
		Description: randomer.RandomOnlyLowerCaseString(30),
		Password:    randomer.RandomString(20),
		Hash:        tokenString,
	}

	//fmt.Println(tokenString)

	// marshal
	suBytes, err := json.Marshal(su)
	require.NoError(t, err)

	// prepare url
	url := "/v1/confirmation/email"

	// prepare request & recorder
	recorder := httptest.NewRecorder()
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(suBytes))
	require.NoError(t, err)

	// send request
	testServer.router.ServeHTTP(recorder, request)

	// parse error
	testParseResponseBody(t, recorder)

	// check status
	require.Less(t, recorder.Code, 300)
}

