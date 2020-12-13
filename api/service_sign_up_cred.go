package api

import (
	"auth/model"
	database "auth/model/sqlc"
	"auth/utils/constants"
	"auth/utils/helper"
	"context"
	"database/sql"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"time"
)

type signUpByEnteringEmail struct {
	Address				string				`json:"address" binding:"email"`
}

// Sign Up
func (server *Server) SignUpSendHashToEmailAddress(ctx *gin.Context) {
	var arg signUpByEnteringEmail
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// check email address
	email, err := server.query.EmailGetOneByAddress(context.Background(), arg.Address)
	if err != nil {
		// pass
	} else if email.ID != 0 {
		ctx.JSON(http.StatusConflict, errorResponse("email address is in use"))
		return
	}

	// create claim
	claim := signUpByEnteringEmailClaim {
		Address:        	arg.Address,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: helper.OnlyGetCurrentTime().Add(time.Hour * 1000).Unix(),
			IssuedAt:  helper.OnlyGetCurrentTime().Unix(),
			Issuer: "auth",
		},
	}

	// get jwt token from claim
	tokenString, err := OnlyGetTokenFromClaim(&claim)
	if err != nil {
		ctx.JSON(http.StatusConflict, errorResponse(err.Error()))
		return
	}

	// prepare a link
	recallLink := url.URL {
		Scheme:     os.Getenv(constants.KEY_BACKEND_SCHEME),
		Host:       os.Getenv(constants.KEY_BACKEND_HOST) + ":" + os.Getenv(constants.KEY_BACKEND_PORT),
		Path: 		"/v1/confirmation/email",
	}

	// set query parameters
	q := recallLink.Query()
	q.Set("hash", tokenString)
	recallLink.RawPath = q.Encode()

	// send token as a link
	notification := model.NotifySignUpConfirmation{
		MailerBasics: model.MailerBasics{
			Language: ctx.GetHeader(constants.KEY_LANGUAGE),
		},
		Address:      arg.Address,
		PreparedLink: recallLink.String(),
	}

	// send message
	model.GetMailerQueue().NotificationChannel <- &notification

	ctx.JSON(http.StatusOK, errorResponse(""))
}

// signUp confirmation
type signUpConfirmHashFromEmail struct {
	Name				string				`json:"name" binding:"require"`
	Description			string				`json:"description"`
	Password			string				`json:"password"`
	Hash				string				`json:"hash" binding:"require"`
}

// confirm hash from email address
func (server *Server) SignUpConfirmHashFromEmail(ctx *gin.Context) {
	// get request body
	var arg signUpConfirmHashFromEmail
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// token -> claim
	claim := signUpByEnteringEmailClaim{}
	status, err := claim.DecodeToken(arg.Hash)
	if err != nil {
		ctx.JSON(status, errorResponse(err.Error()))
		return
	}

	// get user by email address
	user, err := server.query.UserGetOneByEmailAddress(context.Background(), claim.Address)
	if err != nil || user.ID <= 0 {
		ctx.JSON(status, errorResponse("user is not found"))
		return
	}

	// transaction
	tx, err := server.db.BeginTx(context.Background(), nil)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// new queries
	qtx := database.New(tx)
	defer func() {
		if tx != nil {
			_ = tx.Rollback()
		}
	}()

	// create email address
	email, err := qtx.EmailCreateOne(context.Background(), claim.Address)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// create a user
	userCreate := database.UserCreateOneParams{
		Name:        arg.Name,
		Password: 	sql.NullString{
			String: arg.Password,
			Valid:  true,
		},
		Description: sql.NullString{
			String: arg.Description,
			Valid:  true,
		},
		EmailID:     sql.NullInt64{
			Int64: email.ID,
			Valid: true,
		},
		CreatedAt:   sql.NullTime{
			Time:  helper.OnlyGetCurrentTime(),
			Valid: true,
		},
	}

	// create user
	user, err = qtx.UserCreateOne(context.Background(), userCreate)
	if err != nil || user.ID <= 0 {
		ctx.JSON(http.StatusBadRequest, errorResponse("could not create a user"))
		return
	}

	user.Password.String = ""

	ctx.JSON(http.StatusOK, errorResponse(""))
}
