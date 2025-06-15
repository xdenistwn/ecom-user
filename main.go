package main

import (
	"user/cmd/user/handler"
	"user/cmd/user/repository"
	"user/cmd/user/resource"
	"user/cmd/user/service"
	"user/cmd/user/usecase"
	"user/config"
	"user/infrastructure/log"
	"user/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// init config
	cfg := config.LoadConfig()

	// init connection
	redis := resource.InitRedis(&cfg)
	db := resource.InitDb(&cfg)

	// setup logger
	log.SetupLogger()

	// user setup
	userRepository := repository.NewUserRepository(redis, db)
	userService := service.NewUserService(userRepository)
	userUsecase := usecase.NewUserUsecase(userService, cfg.Jwt.Secret)
	userHandler := handler.NewUserHandler(userUsecase)

	port := cfg.App.Port
	router := gin.Default()
	routes.SetupRoutes(router, *userHandler, cfg.Jwt.Secret)

	router.Run(":" + port)

	log.Logger.Printf("Server listening on port: %s", port)
}
