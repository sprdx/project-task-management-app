package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"task-management/lib/databases"
	"task-management/middlewares"
	"task-management/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func CreateUserController(c echo.Context) error {
	var newUser models.NewDataUser
	c.Bind(&newUser)

	// Check if user's data input is valid
	// Tampung return message dari method Validate struct CreateUser
	message := newUser.Validate()
	if message != "OK" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": message,
		})
	}

	// Check if user's email has been registered
	_, count := databases.GetUserByEmail(newUser.Email)
	if count == 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Your email has been registered",
		})
	}

	// Encrypt user's password before insert into database
	hashPassword, _ := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	newUser.Password = string(hashPassword)

	// Exclude ConfirmPassword before insert into database
	user := models.User{Name: newUser.Name, Email: newUser.Email, Password: newUser.Password}

	// Insert value of user struct into database
	_, er := databases.CreateUser(&user)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Failed create user",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Congratulation! User created successfully",
	})
}

func LoginUserController(c echo.Context) error {
	var dataLogin models.LoginUser
	c.Bind(&dataLogin)

	// Check if data of user login is exist and correct in database
	token, err := databases.LoginUser(&dataLogin)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Email or password is incorrect",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Successfull login",
		"data":    token,
	})
}

func GetUserByIDController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Invalid ID",
		})
	}

	// Check if id from token is match to inputted id
	LoggedInId := middlewares.ExtractTokenUserId(c)
	if LoggedInId != id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Access forbidden",
		})
	}

	user, err := databases.GetUserByID(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "User is not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Hello productive people!!",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Invalid ID",
		})
	}

	LoggedInId := middlewares.ExtractTokenUserId(c)
	if LoggedInId != id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Access forbidden",
		})
	}

	user, _ := databases.GetUserByID(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "User is not found",
		})
	}

	var newData models.UpdateDataUser
	c.Bind(&newData)

	message := newData.Validate()
	if message != "OK" {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": message,
		})
	}

	// Check if user email has been registered
	_, count := databases.GetUserByEmail(newData.Email)
	if count == 1 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Inputted email has been registered",
		})
	}

	// Process inputted value of updateDataUser struct into database
	_, er := databases.UpdateUser(id, &newData)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Please input valid password",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Update data user is success",
	})
}

func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Invalid ID",
		})
	}

	LoggedInId := middlewares.ExtractTokenUserId(c)
	if LoggedInId != id {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Access forbidden",
		})
	}

	user, _ := databases.GetUserByID(id)
	if user == nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "User is not found",
		})
	}

	_, er := databases.DeleteUser(id)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Bad request",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Delete user is success",
	})
}

func GetUserByEmailController(c echo.Context) error {
	email := c.Param("email")
	fmt.Println(email)
	us, count := databases.GetUserByEmail(email)

	if count == 0 {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Data not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": us,
	})
}
