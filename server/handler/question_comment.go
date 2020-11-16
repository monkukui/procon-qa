package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/server/entity"
)

// qid からコメントを全取得する
func GetQuestionComments(c echo.Context) error {

	qid, err := strconv.Atoi(c.Param("qid"))
	fmt.Println("hogehoge")
	if err != nil {
		return echo.ErrNotFound
	}

	comments := entity.FindQuestionComments(&entity.QuestionComment{QID: qid})
	return c.JSON(http.StatusOK, comments)
}

func PostQuestionComment(c echo.Context) error {
	comment := new(entity.QuestionComment)
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
	if user := entity.FindUser(&entity.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	comment.UID = uid

	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	nowJST := nowUTC.In(jst)
	comment.Date = nowJST.Format("2006/01/02 15:04:05")

	entity.CreateQuestionComment(comment)

	return c.JSON(http.StatusCreated, comment)
}

func CountQuestionComment(c echo.Context) error {
	return c.JSON(http.StatusOK, len(entity.FindQuestionComments(&entity.QuestionComment{})))
}

func DeleteQuestionComment(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := entity.FindUser(&entity.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	// uid が question の uid と一致していなければダメ
	if uid != entity.FindQuestionComments(&entity.QuestionComment{ID: id})[0].UID {
		return echo.ErrNotFound
	}
	if err := entity.DeleteQuestionComment(&entity.QuestionComment{ID: id, UID: uid}); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}
