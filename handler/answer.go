package handler

import (
	"net/http"
	"strconv"
  //"fmt"
	"github.com/labstack/echo"
	// c-color さんに依存しちゃってる...
	// なんとか再現性があるように作り変えることができないもんかね...?
	// でも, wifi なくても起動できるんだよなぁ
	// どこを参照しているんだろう.....
	// "github.com/x-color/simple-webapp/model"
	// ? できたのか ?
	"github.com/monkukui/procon-qa/model"
)

// 質問に紐ずいた, 回答を全取得する
func GetAnswersForQuestion(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	QuestionID, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		return echo.ErrNotFound
	}

	answers := model.FindAnswers(&model.Answer{QID: QuestionID})
	return c.JSON(http.StatusOK, answers)
}

// 回答を 1 つ 取得する
func GetAnswer(c echo.Context) error {

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	AnswerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	answer := model.FindAnswers(&model.Answer{ID: AnswerID})[0]
	return c.JSON(http.StatusOK, answer)
}

// 回答を投稿する
func PostAnswer(c echo.Context) error {
	answer := new(model.Answer)

	// answer に 送信されてきたデータを bind している
	if err := c.Bind(answer); err != nil {
		return err
	}

	// 妥当性判定
	// Body が空欄ではないことをチェックする
	if answer.Body == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  // 回答者をユーザーに設定
	answer.UID = uid
	model.CreateAnswer(answer)

	return c.JSON(http.StatusCreated, answer)
}

