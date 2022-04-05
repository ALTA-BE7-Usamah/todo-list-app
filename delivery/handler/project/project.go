package project

import (
	"net/http"
	"project2/todo-list-app/delivery/helper"
	_middlewares "project2/todo-list-app/delivery/middlewares"
	"project2/todo-list-app/entities"
	_projectUseCase "project2/todo-list-app/usecase/project"
	"strconv"

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

func (ph *ProjectHandler) GetAllProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		projects, err := ph.projectUseCase.GetAllProject(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseProject := []map[string]interface{}{}
		for i := 0; i < len(projects); i++ {
			response := map[string]interface{}{
				"ID":           projects[i].ID,
				"project_name": projects[i].ProjectName,
			}
			responseProject = append(responseProject, response)
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all project", responseProject))
	}
}

func (ph *ProjectHandler) GetProjectbyIdHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		project, rows, err := ph.projectUseCase.GetProjectbyId(uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}

		responseProject := map[string]interface{}{
			"ID":           project.ID,
			"project_name": project.ProjectName,
			"task":         project.Task,
		}

		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get project", responseProject))
	}
}

func (ph *ProjectHandler) AddTaskProjectHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		var addTask entities.Task
		errBind := c.Bind(&addTask)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		_, rows, err := ph.projectUseCase.AddTaskProject(addTask, uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success add task to project"))
	}
}
