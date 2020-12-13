package api

import (
	"context"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type signInByEnteringEmail struct {
	Address  string       `json:"address"`
	Password  string       `json:"password"`
}

func (server *Server) SignInByEnteringEmail(ctx *gin.Context) {
	var arg signInByEnteringEmail
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// logic goes here
	email, err := server.query.EmailGetOneByAddress(context.Background(), arg.Address)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err.Error()))
		return
	}

	// get user by email_id
	user, err := server.query.UserGetOneByEmailAddress(context.Background(), email.Address)
	if err != nil || user.ID <= 0 {
		ctx.JSON(http.StatusNotFound, errorResponse("user is not found"))
		return
	} else if user.Password.Valid == false {
		ctx.JSON(http.StatusNotFound, errorResponse("please, enter using google sign in"))
		return
	}

	// compare password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password.String), []byte(arg.Password))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		ctx.JSON(http.StatusBadRequest, errorResponse("password is invalid"))
		return
	} else if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// email address & password are valid
	user.Password.String = ""



	ctx.JSON(http.StatusOK, user)
}
