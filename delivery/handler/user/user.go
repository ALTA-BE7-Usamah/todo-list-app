package user

import (
	"net/http"
	"project2/todo-list-app/delivery/helper"
	_middlewares "project2/todo-list-app/delivery/middlewares"
	"project2/todo-list-app/entities"
	_userUseCase "project2/todo-list-app/usecase/user"
	"strconv"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userUseCase _userUseCase.UserUseCaseInterface
}

func NewUserHandler(userUseCase _userUseCase.UserUseCaseInterface) *UserHandler {
	return &UserHandler{
		userUseCase: userUseCase,
	}
}

func (uh *UserHandler) CreateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		var newUser entities.User
		err := c.Bind(&newUser)
		if err != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}
		_, error := uh.userUseCase.CreatUser(newUser)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success create user"))
	}
}

func (uh *UserHandler) GetUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		// check apakah id dari token sama dengan id dari parm
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		//jika id sama dan tidak ada error
		user, rows, err := uh.userUseCase.GetUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseUser := map[string]interface{}{
			"ID":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"gender":     user.Gender,
			"age":        user.Age,
			"profession": user.Profession,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success get user", responseUser))
	}
}

func (uh *UserHandler) UpdateUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		// check apakah id dari token sama dengan id dari parm
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}

		//jika id sama dan tidak ada error
		var updateUser entities.User
		errBind := c.Bind(&updateUser)
		if errBind != nil {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("failed to bind data. please check your data"))
		}

		user, rows, error := uh.userUseCase.UpdateUser(updateUser, id)
		if error != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		responseUser := map[string]interface{}{
			"ID":         user.ID,
			"name":       user.Name,
			"email":      user.Email,
			"gender":     user.Gender,
			"age":        user.Age,
			"profession": user.Profession,
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccess("success update user", responseUser))
	}
}

func (uh *UserHandler) DeleteUserHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		//mendapatkan id dari token yang dimasukkan
		idToken, errToken := _middlewares.ExtractToken(c)
		if errToken != nil {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}

		idStr := c.Param("id")
		id, err := strconv.Atoi(idStr)
		// check apakah id dari token sama dengan id dari parm
		if idToken != id {
			return c.JSON(http.StatusUnauthorized, helper.ResponseFailed("unauthorized"))
		}
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("id not recognise"))
		}
		//jika id sama dan tidak ada error
		rows, err := uh.userUseCase.DeleteUser(id)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, helper.ResponseFailed("failed to fetch data"))
		}
		if rows == 0 {
			return c.JSON(http.StatusBadRequest, helper.ResponseFailed("data not found"))
		}
		return c.JSON(http.StatusOK, helper.ResponseSuccessWithoutData("success deleted user"))
	}
}
