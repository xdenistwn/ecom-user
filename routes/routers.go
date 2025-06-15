package routes

import (
	"user/cmd/user/handler"
	"user/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine, userHandler handler.UserHandler, jwtSecret string) {
	// context timeout and logger
	router.Use(middleware.RequestLogger(2))

	// health check
	router.GET("/ping", userHandler.Ping)

	// Public API
	public := router.Group("/api/v1/auth")
	public.POST("/register", userHandler.Register)
	public.POST("/login", userHandler.Login)

	// Private API
	authMiddleware := middleware.AuthMiddleware(jwtSecret)
	private := router.Group("/api/v1/user/")
	private.Use(authMiddleware)
	private.GET("/user-info", userHandler.GetUserInfo)
}
