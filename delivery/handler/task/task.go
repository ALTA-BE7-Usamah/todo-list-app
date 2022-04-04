package task

import (
	"net/http"
	"project2/todo-list-app/delivery/helper"
	_middlewares "project2/todo-list-app/delivery/middlewares"
	"project2/todo-list-app/entities"
	_taskUseCase "project2/todo-list-app/usecase/task"
	"strconv"

	"github.com/labstack/echo/v4"
)

type TaskHandler struct {
	taskUseCase _taskUseCase.TaskUseCaseInterface
}

func NewTaskHandler(taskUseCase _taskUseCase.TaskUseCaseInterface) *TaskHandler {
	return &TaskHandler{
		taskUseCase: taskUseCase,
	}
}

func (th *TaskHandler) CreateTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		var newTask entities.Task
		err := c.Bind(&newTask)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		// memastikan bahwa yang membuat project adalah yang login
		newTask.UserID = uint(idToken)

		_, error := th.taskUseCase.CreatTask(newTask)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create task"))
	}
}

func (th *TaskHandler) GetAllTaskHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		tasks, err := th.taskUseCase.GetAllTask(uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}

		responseTask := []map[string]interface{}{}
		for i := 0; i < len(tasks); i++ {
			response := map[string]interface{}{
				"ID":          tasks[i].ID,
				"task_name":   tasks[i].TaskName,
				"task_status": tasks[i].TaskStatus,
			}
			responseTask = append(responseTask, response)
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get all tasks", responseTask))
	}
}

func (th *TaskHandler) GetTaskByIdIDHandler() echo.HandlerFunc {
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
		task, rows, err := th.taskUseCase.GetTaskById(uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get rent", task))
	}
}

func (th *TaskHandler) UpdateTaskHandler() echo.HandlerFunc {
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

		var updateTask entities.Task
		errBind := c.Bind(&updateTask)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		task, rows, err := th.taskUseCase.UpdateTask(updateTask, uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseTask := map[string]interface{}{
			"ID":          task.ID,
			"task_name":   task.TaskName,
			"task_status": task.TaskStatus,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get rent", responseTask))
	}
}

func (th *TaskHandler) DeleteTaskHandler() echo.HandlerFunc {
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
		rows, err := th.taskUseCase.DeleteTask(uint(id), uint(idToken))
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("unauthorized"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success deleted task"))
	}
}
