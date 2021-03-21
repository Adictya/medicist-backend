package api

import (
	"net/http"

	db "github.com/adictya/medicist-backend/db/sqlc"
	"github.com/gin-gonic/gin"
)

type createAccountRequest struct {
	FullName string `json:"FullName" binding:"required" `
	Mobile   int64  `json:"Mobile" binding:"required" `
}

func(server *Server) createAccount(ctx *gin.Context){
	var req createAccountRequest
	if err:=  ctx.ShouldBindJSON(&req); err != nil{
		ctx.JSON(http.StatusBadRequest,errorResponse(err))
		return
	}

	arg := db.CreateAccountParams{
		FullName: req.FullName,
		Mobile: req.Mobile,
		Type: 0,
	}

	account, err := server.store.CreateAccount(ctx,arg)
	if err!=nil{
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK,account)

}
