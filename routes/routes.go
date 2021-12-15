package routes

import (
	"task-management/constants"
	"task-management/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/users", controllers.CreateUserController)

	e.GET("/users/:email", controllers.GetUserByEmailController)
	e.POST("/login", controllers.LoginUserController)

	e.POST("/tasks", controllers.CreateTaskController)

	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))
	r.GET("/users/:id", controllers.GetUserByIDController)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

	return e
}
