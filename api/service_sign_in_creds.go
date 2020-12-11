package api

import (
	database "auth/model/sqlc"
	"auth/utils/helper"
	"database/sql"
	"github.com/gin-gonic/gin"
	"net/http"
)

type signInWithCredentials struct {
	Username  string       `json:"username"`
	Password  string       `json:"password"`
	CreatedAt sql.NullTime `json:"created_at"`
}

func (server *Server) createBoard(ctx *gin.Context) {
	var arg signInWithCredentials
	if err := ctx.ShouldBindJSON(&arg); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err.Error()))
		return
	}

	// set current time
	arg.CreatedAt = sql.NullTime{
		Time:  helper.OnlyGetCurrentTime(),
		Valid: true,
	}

	// logic goes here
	board, err := server.query.CredentialCreateOne(ctx, database.CredentialCreateOneParams(arg))
	if err != nil {
		errorResponse(err.Error())
	}

	ctx.JSON(http.StatusOK, board)
}
