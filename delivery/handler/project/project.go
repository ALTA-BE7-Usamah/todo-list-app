package project

import (
	"net/http"
	"project2/todo-list-app/delivery/helper"
	_middlewares "project2/todo-list-app/delivery/middlewares"
	"project2/todo-list-app/entities"
	_projectUseCase "project2/todo-list-app/usecase/project"

	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	projectUseCase _projectUseCase.ProjectUseCaseInterface
}

func NewProjectHandler(projectUseCase _projectUseCase.ProjectUseCaseInterface) *ProjectHandler {
	return &ProjectHandler{
		projectUseCase: projectUseCase,
	}
}

func (ph *ProjectHandler) CreateProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newproject entities.Project
		err := c.Bind(&newproject)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		// memastikan bahwa yang membuat project adalah yang login
		newproject.UserID = uint(idToken)

		_, error := ph.projectUseCase.CreatProject(newproject)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create project"))
	}
}
