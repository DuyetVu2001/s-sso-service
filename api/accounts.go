package api

import (
	"net/http"
	db "sso-service/db/sqlc"
	"sso-service/util"
	"time"

	"github.com/gin-gonic/gin"
)

type registerRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

type userResponse struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Email     *string   `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

func (server *Server) register(ctx *gin.Context) {
	var req registerRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	args := db.CreateAccountParams{
		Username:     req.Username,
		PasswordHash: &hashedPassword,
	}

	user, err := server.store.CreateAccount(ctx, args)
	if err != nil {
		if db.ErrorCode(err) == db.ErrUniqueViolation.Code {
			ctx.JSON(http.StatusForbidden, errorResponse(err))
			return
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"user": user})
}

type loginRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.store.GetAccountByUsername(ctx, req.Username)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}

	err = util.VerifyPassword(*user.PasswordHash, req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	res := userResponse{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}

	ctx.JSON(http.StatusOK, gin.H{"status": "you are logged in", "user": res})
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
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"account": account})
}

type getListAccountsRequest struct {
	Offset int32 `form:"offset" binding:"required,min=1"`
	Limit  int32 `form:"limit" binding:"required,min=1,max=50"`
}

func (server *Server) getListAccounts(ctx *gin.Context) {
	var req getListAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	args := db.GetListAccountsParams{
		Limit:  req.Limit,
		Offset: (req.Offset - 1) * req.Limit,
	}

	accounts, err := server.store.GetListAccounts(ctx, args)
	if err != nil {

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"accounts": accounts})
}
