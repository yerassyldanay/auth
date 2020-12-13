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

func NewServer(query database.Querier) *Server {
	server := &Server{
		query: query,
	}
	router := gin.Default()

	//router.POST("/board", server.S)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(errmsg string) gin.H {
	return gin.H { "error": errmsg }
}
