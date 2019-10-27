package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/monkukui/procon-qa/handler"
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

  api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))
	// ここから変更した

	// questions
	api.GET("/questions", handler.GetAllQuestions)            // 質問の全取得
	api.GET("/questions/:page", handler.GetQuestionsWithPage) // 質問をページ全取得
	api.GET("/user-questions/:page", handler.GetUserQuestionsWithPage) // 質問を 1 つ取得
	api.GET("/question/:id", handler.GetQuestion)             // 質問を 1 つ取得
	api.POST("/questions", handler.PostQuestion)              // 質問の投稿
	api.DELETE("/question/:id", handler.DeleteQuestion)       // 質問の削除
	api.PUT("/question/:id/completed", handler.UpdateQuestionCompleted)  // 質問の更新

  // answers
	api.GET("/answers/:qid", handler.GetAnswersForQuestion)   // 質問に紐づいた 回答を全取得
	api.GET("/answer/:id", handler.GetAnswer)                 // 回答を 1 つ取得
	api.GET("/user-answers/:page", handler.GetUserAnswersWithPage) // 質問を 1 つ取得
	api.POST("/answers", handler.PostAnswer)    // 回答の投稿
  // users
	api.GET("/user/:uid", handler.GetUser) // user_id から ユーザー名を取得
	return e
}
