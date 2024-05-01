package gapi

import (
	db "sso-service/db/sqlc"
	"sso-service/pb"

	"github.com/gin-gonic/gin"
)

type Server struct {
	pb.UnimplementedAppServiceServer
	store  *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server {
	server := &Server{store: store}
	return server
}

func (server *Server) Start(address string) error {

	return server.router.Run(address)
}
