package api

import (
	"fmt"

	db "github.com/QuyenFunc/Goaudio/db/sqlc"
	"github.com/QuyenFunc/Goaudio/db/util"
	"github.com/QuyenFunc/Goaudio/token"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// if v, ok := binding.Validator.Engine().(*Validator.Validate); ok {
	// 	v.RegisterValidation("cu")
	// }

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))
	// router.POST("/catalogs", server.createCatalog)
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	router.POST("/products", server.createProduct)
	router.GET("/products/:id", server.getProduct)
	router.GET("/products", server.listProduct)
	server.router = router
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
