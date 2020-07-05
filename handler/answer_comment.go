package handler

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

// aid からコメントを全取得する
func GetAnswerComments(c echo.Context) error {

	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		return echo.ErrNotFound
	}

	comments := model.FindAnswerComments(&model.AnswerComment{AID: aid})
	return c.JSON(http.StatusOK, comments)
}

func PostAnswerComment(c echo.Context) error {
	comment := new(model.AnswerComment)
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

	model.CreateAnswerComment(comment)

	return c.JSON(http.StatusCreated, comment)
}

func CountAnswerComment(c echo.Context) error {
	return c.JSON(http.StatusOK, len(model.FindAnswerComments(&model.AnswerComment{})))
}
