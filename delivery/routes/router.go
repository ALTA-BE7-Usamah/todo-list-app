package routes

import (
	_authHandler "project2/todo-list-app/delivery/handler/auth"
	_projectHandler "project2/todo-list-app/delivery/handler/project"
	_taskHandler "project2/todo-list-app/delivery/handler/task"
	_userHandler "project2/todo-list-app/delivery/handler/user"
	_middlewares "project2/todo-list-app/delivery/middlewares"

	"github.com/labstack/echo/v4"
)

func RegisterAuthPath(e *echo.Echo, ah *_authHandler.AuthHandler) {
	e.POST("/auth", ah.LoginHandler())
}

func RegisterUserPath(e *echo.Echo, uh *_userHandler.UserHandler) {
	e.POST("/users", uh.CreateUserHandler())
	e.GET("/users/:id", uh.GetUserHandler(), _middlewares.JWTMiddleware())
	e.PUT("/users/:id", uh.UpdateUserHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/users/:id", uh.DeleteUserHandler(), _middlewares.JWTMiddleware())
}

func RegisterProjectPath(e *echo.Echo, ph *_projectHandler.ProjectHandler) {
	e.POST("/projects", ph.CreateProjectHandler(), _middlewares.JWTMiddleware())
	e.GET("/projects", ph.GetAllProjectHandler(), _middlewares.JWTMiddleware())
	e.GET("/projects/:id", ph.GetProjectbyIdHandler(), _middlewares.JWTMiddleware())
}

func RegisterTaskPath(e *echo.Echo, th *_taskHandler.TaskHandler) {
	e.POST("/todo/tasks", th.CreateTaskHandler(), _middlewares.JWTMiddleware())
	e.GET("/todo/tasks", th.GetAllTaskHandler(), _middlewares.JWTMiddleware())
	e.GET("/todo/tasks/:id", th.GetTaskByIdIDHandler(), _middlewares.JWTMiddleware())
	e.PUT("/todo/tasks/:id", th.UpdateTaskHandler(), _middlewares.JWTMiddleware())
	e.DELETE("/todo/tasks/:id", th.DeleteTaskHandler(), _middlewares.JWTMiddleware())
}
