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
	e.Static("/fonts", "client/dist/fonts")
	e.Static("/img", "client/dist/img")
	e.File("/", "client/dist/index.html")
	e.File("/signup", "public/signup.html")
	e.POST("/signup", handler.Signup)
	e.File("/login", "public/login.html")
	e.POST("/login", handler.Login)

	api := e.Group("/api")
	api.Use(middleware.JWTWithConfig(handler.Config))

	// 認証なしで呼べる api たち
	noAuth := e.Group("/api/no-auth")
	noAuth.GET("/questions", handler.GetAllQuestions)                          // 質問の全取得
	noAuth.GET("/questions/count", handler.GetQuestionSize)                    // 質問の個数を取得
	noAuth.GET("/completed-questions/count", handler.GetCompletedQuestionSize) // 解決済みの質問の個数を取得
	noAuth.GET("/answers/count", handler.GetAnswerSize)                        // 回答の個数を取得
	noAuth.GET("/users/count", handler.GetUserSize)                            // ユーザの個数を取得
	noAuth.GET("/questions/:page/:mode", handler.GetQuestionsWithPage)         // 質問をページ取得(mode あり)
	noAuth.GET("/users/:page/:mode", handler.GetUsersWithPage)                 // ユーザをページ取得(mode あり)
	noAuth.GET("/question/:id", handler.GetQuestion)                           // 質問を 1 つ取得
	noAuth.GET("/answers/:qid/:mode", handler.GetAnswersForQuestion)           // 質問に紐づいた 回答を全取得
	noAuth.GET("/answer/:id", handler.GetAnswer)                               // 回答を 1 つ取得
	noAuth.GET("/user/:uid", handler.GetUser)                                  // user_id から ユーザー名を取得
	noAuth.PUT("/question/:id/browse", handler.BrowseQuestion)                 // 閲覧

	// questions
	api.GET("/questions", handler.GetAllQuestions)                      // 質問の全取得
	api.GET("/questions/count", handler.GetQuestionSize)                // 質問の個数を取得
	api.GET("/questions/:page", handler.GetQuestionsWithPage)           // 質問をページ全取得
	api.GET("/user-questions/:page", handler.GetUserQuestionsWithPage)  // 質問を 1 つ取得
	api.GET("/question/:id", handler.GetQuestion)                       // 質問を 1 つ取得
	api.POST("/questions", handler.PostQuestion)                        // 質問の投稿
	api.DELETE("/question/:id", handler.DeleteQuestion)                 // 質問の削除
	api.PUT("/question/:id/completed", handler.UpdateQuestionCompleted) // 質問の完了フラグの更新
	api.PUT("/question/:id/favorite", handler.FavoriteQuestion)         // いいね

	// answers
	api.GET("/answers/:qid", handler.GetAnswersForQuestion)        // 質問に紐づいた 回答を全取得
	api.GET("/answer/:id", handler.GetAnswer)                      // 回答を 1 つ取得
	api.GET("/user-answers/:page", handler.GetUserAnswersWithPage) // 質問を 1 つ取得
	api.POST("/answers", handler.PostAnswer)                       // 回答の投稿
	api.DELETE("/answer/:id", handler.DeleteAnswer)                // 質問の削除
	api.PUT("/answer/:id/favorite", handler.FavoriteAnswer)        // いいね
	// users
	api.GET("/user/:uid", handler.GetUser)       // user_id から ユーザー名を取得
	api.DELETE("/user/:uid", handler.DeleteUser) // user_id から ユーザー名を取得

	// question_good
	api.GET("/question-good/:uid/:qid", handler.QuestionFavorited) // いいね状態かどうか
	// answer_good
	api.GET("/answer-good/:uid/:aid", handler.AnswerFavorited) // いいね状態かどうか

	api.GET("/token", handler.Token)

	return e
}
