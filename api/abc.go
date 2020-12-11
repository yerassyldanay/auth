package api

import (
	database "auth/model/sqlc"
	"github.com/gin-gonic/gin"
)

type Server struct {
	query		database.Querier
	router		*gin.Engine
}

func NewServer(query database.Querier) *Server {
	server := &Server{
		query: query,
	}
	router := gin.Default()

	router.POST("/board", server.createBoard)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(errmsg string) gin.H {
	return gin.H { "error": errmsg }
}