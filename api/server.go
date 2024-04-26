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

	router.GET("/accounts", func(c *gin.Context) {})
	router.POST("/accounts", func(c *gin.Context) {})
	router.PUT("/accounts/change-password", func(c *gin.Context) {})

	server.router = router
	return server
}
