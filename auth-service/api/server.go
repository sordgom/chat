package api

import (
	db "chat-app/auth-service/db/sqlc"
	"chat-app/auth-service/token"
	"chat-app/util"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	config     util.Config
	auth       db.Auth
	tokenMaker token.Maker
	router     *gin.Engine
}

func NewServer(config util.Config, auth db.Auth) (*Server, error) {
	tokenMaker, err := token.NewJWTMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}
	server := &Server{
		config:     config,
		auth:       auth,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()

	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
	}))

	auth := router.Group("/auth")
	{
		auth.POST("", server.createUser)
		auth.POST("/login", server.loginUser)
		auth.POST("/token", server.createAccessToken)
		auth.POST("/refresh", server.renewAccessToken)
		auth.Use(authMiddleware(server.tokenMaker))
		{
			auth.GET("/users/:id", server.getUser)
			auth.GET("/users", server.listUsers)
		}
	}

	server.router = router
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
