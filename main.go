package main

import (
	"fmt"
	"log"
	"project2/todo-list-app/configs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	_authHandler "project2/todo-list-app/delivery/handler/auth"
	_authRepository "project2/todo-list-app/repository/auth"
	_authUseCase "project2/todo-list-app/usecase/auth"

	_userHandler "project2/todo-list-app/delivery/handler/user"
	_userRepository "project2/todo-list-app/repository/user"
	_userUseCase "project2/todo-list-app/usecase/user"

	_middlewares "project2/todo-list-app/delivery/middlewares"
	_routes "project2/todo-list-app/delivery/routes"
	_utils "project2/todo-list-app/utils"
)

func main() {
	config := configs.GetConfig()
	db := _utils.InitDB(config)

	authRepo := _authRepository.NewAuthRepository(db)
	authUseCase := _authUseCase.NewAuthUseCase(authRepo)
	authHandler := _authHandler.NewAuthHandler(authUseCase)

	userRepo := _userRepository.NewUserRepository(db)
	userUseCase := _userUseCase.NewUserUseCase(userRepo)
	userHandler := _userHandler.NewUserHandler(userUseCase)

	e := echo.New()
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(_middlewares.CustomLogger())

	_routes.RegisterAuthPath(e, authHandler)
	_routes.RegisterUserPath(e, userHandler)

	log.Fatal(e.Start(fmt.Sprintf(":%v", config.Port)))
}
