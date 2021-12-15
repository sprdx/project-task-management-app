package controllers

import (
	"fmt"
	"net/http"
	"task-management/lib/databases"
	"task-management/models"

	"github.com/labstack/echo/v4"
)

func CreateTaskController(c echo.Context) error {
	var newTask models.Task
	c.Bind(&newTask)
	fmt.Println(newTask)

	_, er := databases.CreateTask(&newTask)
	if er != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "Bad request",
			"message": "Failed create task",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "Success",
		"message": "Congratulation! Task created successfully",
	})
}
