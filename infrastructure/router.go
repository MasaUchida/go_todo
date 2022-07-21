package infrastructure

import (
	"example/todo/interfaces/controller"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Init() {
	e := echo.New()

	taskController := controller.NewTaskController(NewSqlHandler())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/tasks", func(ctx echo.Context) error { return taskController.Index(ctx) })
	e.GET("/task/:id", func(ctx echo.Context) error { return taskController.Show(ctx) })
	e.POST("/create", func(ctx echo.Context) error { return taskController.Create(ctx) })

	e.Logger.Fatal(e.Start(":1323"))
}
