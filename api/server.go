package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"

	db "github.com/wizlif/simplebank/db/sqlc"
	"github.com/wizlif/simplebank/token"
	"github.com/wizlif/simplebank/util"
)

// Server serves HTTP requests for our banking requests
type Server struct {
	config     util.Config
	db         db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)

	if err != nil {
		return nil, fmt.Errorf("cannot create token: %w", err)
	}

	server := &Server{
		config:     config,
		db:         store,
		tokenMaker: tokenMaker,
	}

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("currency", validCurrency)
	}

	server.setUpRouter()

	return server, nil
}

// setup router
func (server *Server) setUpRouter() {
	router := gin.Default()

	router.POST("/users", server.CreateUser)
	router.POST("/users/login", server.LoginUser)
	router.POST("/tokens/renew_access", server.renewAccessToken)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.CreateAccount)
	authRoutes.GET("/accounts/:id", server.GetAccount)
	authRoutes.GET("/accounts", server.ListAccounts)
	authRoutes.PATCH("/accounts/:id", server.UpdateAccount)
	authRoutes.DELETE("/accounts/:id", server.DeleteAccount)

	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a particular address
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
