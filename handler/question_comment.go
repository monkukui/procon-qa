package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

// qid からコメントを全取得する
func GetQuestionComments(c echo.Context) error {

	qid, err := strconv.Atoi(c.Param("qid"))
	fmt.Println("hogehoge")
	if err != nil {
		return echo.ErrNotFound
	}

	comments := model.FindQuestionComments(&model.QuestionComment{QID: qid})
	return c.JSON(http.StatusOK, comments)
}

func PostQuestionComment(c echo.Context) error {
	comment := new(model.QuestionComment)
	if err := c.Bind(comment); err != nil {
		return err
	}

	// 妥当性判定
	// Body が空欄ではないことをチェック
	if comment.Body == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	comment.UID = uid

	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	comment.Date = nowJST.Format("2006/01/02 15:04:05")

	model.CreateQuestionComment(comment)

	return c.JSON(http.StatusCreated, comment)
}

func CountQuestionComment(c echo.Context) error {
	return c.JSON(http.StatusOK, len(model.FindQuestionComments(&model.QuestionComment{})))
}

func DeleteQuestionComment(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	// uid が question の uid と一致していなければダメ
	if uid != model.FindQuestionComments(&model.QuestionComment{ID: id})[0].UID {
		return echo.ErrNotFound
	}
	if err := model.DeleteQuestionComment(&model.QuestionComment{ID: id, UID: uid}); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}
