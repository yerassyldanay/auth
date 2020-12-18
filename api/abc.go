package api

import (
	database "auth/model/sqlc"
	"database/sql"
	"github.com/gin-gonic/gin"
)

type Server struct {
	db			*sql.DB
	query		database.Querier
	router		*gin.Engine
}

func NewServer(connection *sql.DB) *Server {
	gin.SetMode(gin.TestMode)

	server := &Server{
		query: database.New(connection),
		db: connection,
	}
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(gin.Logger())

	v1 := router.Group("/v1")

	// sign
	v1.POST("/sign-up/email", server.SignUpSendHashToEmailAddress)

	// confirm
	v1.POST("/confirmation/email", server.SignUpConfirmHashFromEmail)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(errmsg string) gin.H {
	return gin.H { "error": errmsg }
}
