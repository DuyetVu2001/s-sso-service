package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type loginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Username != "manu" || req.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if req.Username != "manu" || req.Password != "123" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
}

type getAccountByIdRequest struct {
	ID int64 `uri:"id" biding:"required,min=1"`
}

func (server *Server) getAccountById(ctx *gin.Context) {
	var req getAccountByIdRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	account, err := server.store.GetAccountById(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
	}

	ctx.JSON(http.StatusOK, gin.H{"account": account})
}
