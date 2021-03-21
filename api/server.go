package api

import (
	db "github.com/adictya/medicist-backend/db/sqlc"
	"github.com/gin-gonic/gin"

)

//server will serve all http request
type Server struct{
	store *db.Store
	router *gin.Engine
}

func NewServer(store *db.Store) *Server{
	server := &Server{ store: store }
	router := gin.Default()

	router.POST("/accounts",server.createAccount)

	server.router = router
	return server
}

func (server *Server) Start(address string) error{
	return server.router.Run(address)
}

func errorResponse(err error) gin.H{
	return gin.H{ "Error":err.Error() }
}


