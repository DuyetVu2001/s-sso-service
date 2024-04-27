package api

import (
	db "sso-service/db/sqlc"

	"github.com/gin-gonic/gin"
)

type Server struct {
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	router.GET("/accounts", server.getListAccounts)
	router.GET("/accounts/:id", server.getAccountById)

	router.POST("/auth/register", server.register)
	router.POST("/auth/login", server.login)
	router.PUT("/auth/change-password", func(c *gin.Context) {})

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
