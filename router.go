package main

import (
    "github.com/labstack/echo"
    "github.com/labstack/echo/middleware"
    "github.com/x-color/simple-webapp/handler"
)

func newRouter() *echo.Echo {
	e := echo.New()

    e.Use(middleware.Logger())
    e.Use(middleware.Recover())

    e.Static("/js", "client/dist/js")
    e.Static("/css", "client/dist/css")

    e.File("/", "client/dist/index.html")
    e.File("/signup", "public/signup.html")
    e.POST("/signup", handler.Signup)
    e.File("/login", "public/login.html")
    e.POST("/login", handler.Login)
    e.File("/todos", "public/todos.html")

    api := e.Group("/api")
    api.Use(middleware.JWTWithConfig(handler.Config))
    api.GET("/todos", handler.GetTodos)
    api.POST("/todos", handler.AddTodo)
    api.DELETE("/todos/:id", handler.DeleteTodo)
    api.PUT("/todos/:id/completed", handler.UpdateTodo)

    // ここから変更した
    api.GET("/questions", handler.GetQuestions)

	return e
}
