package main

import (
	"net"
	"user/cmd/user/handler"
	"user/cmd/user/repository"
	"user/cmd/user/resource"
	"user/cmd/user/service"
	"user/cmd/user/usecase"
	"user/config"
	"user/infrastructure/log"
	"user/proto/userpb"
	"user/routes"

	grpcUser "user/grpc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
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

	go func() {
		// HTTP server setup
		port := cfg.App.Port
		router := gin.Default()
		routes.SetupRoutes(router, *userHandler, cfg.Jwt.Secret)
		router.Run(":" + port)
		log.Logger.Printf("HTTP Server listening on port: %s", port)
	}()

	listen, err := net.Listen("tcp", ":"+cfg.App.GRPCPort)
	if err != nil {
		log.Logger.Fatalf("Failed to listen on port %s: %v", cfg.App.GRPCPort, err)
	}

	// GRPC server setup
	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, &grpcUser.GRPCServer{
		UserUsecase: *userUsecase,
	})
	reflection.Register(grpcServer)
	log.Logger.Printf("GRPC Server listening on port: %s", cfg.App.GRPCPort)
	if err := grpcServer.Serve(listen); err != nil {
		log.Logger.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
