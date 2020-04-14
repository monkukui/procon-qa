package handler

import (
	"github.com/labstack/echo"
	"net/http"
	"strconv"
	"time"
	// c-color さんに依存しちゃってる...
	// なんとか再現性があるように作り変えることができないもんかね...?
	// でも, wifi なくても起動できるんだよなぁ
	// どこを参照しているんだろう.....
	// "github.com/x-color/simple-webapp/model"
	// ? できたのか ?
	"github.com/monkukui/procon-qa/model"
)

// answers のサイズを取得する
func GetAnswerSize(c echo.Context) error {
	return c.JSON(http.StatusOK, len(model.FindAnswers(&model.Answer{}, "id desc")))
}

// 質問に紐ずいた, 回答を全取得する
func GetAnswersForQuestion(c echo.Context) error {
	// uid := userIDFromToken(c)
	// if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
	//	return echo.ErrNotFound
	// }

	QuestionID, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		return echo.ErrNotFound
	}
	modeId, err := strconv.Atoi(c.Param("mode")) // ソートの設定
	if err != nil {
		return echo.ErrNotFound
	}

	mode := "id desc"

	if modeId == 2 {
		mode = "favorite_count desc"
	}

	answers := model.FindAnswers(&model.Answer{QID: QuestionID}, mode)
	return c.JSON(http.StatusOK, answers)
}

// userId/pageId で質問をページ取得する
func GetUserAnswersWithPage(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	PageLength := 5                              // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}

	answers := model.FindAnswersWithPage(&model.Answer{UID: uid}, PageID, PageLength)
	return c.JSON(http.StatusOK, answers)
}

// 回答を 1 つ 取得する
func GetAnswer(c echo.Context) error {

	// uid := userIDFromToken(c)
	// if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
	// 	return echo.ErrNotFound
	// }

	AnswerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	answer := model.FindAnswers(&model.Answer{ID: AnswerID}, "id desc")[0]
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
	answer.FavoriteCount = 0
	answer.Date = time.Now().Format("2006/01/02 15:04:05")

	model.CreateAnswer(answer)

	return c.JSON(http.StatusCreated, answer)
}

// 回答を削除する
func DeleteAnswer(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	answerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	if err := model.DeleteAnswer(&model.FindAnswers(&model.Answer{ID: answerID}, "id desc")[0]); err != nil {
    return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

// いいねをする
func FavoriteAnswer(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	answerID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	answers := model.FindAnswers(&model.Answer{ID: answerID}, "id desc")
	if len(answers) == 0 {
		return echo.ErrNotFound
	}
	answer := answers[0]

	if answer.UID == uid {
		// 自分の回答にいいねはできません
		return echo.ErrNotFound
	}

	// question と同じロジック
	goods := model.FindAnswerGoods(&model.AnswerGood{UID: uid, AID: answerID})

	if len(goods) == 0 { // いいねをする
		model.CreateAnswerGood(&model.AnswerGood{UID: uid, AID: answerID})
		answer.FavoriteCount++
		if err := model.UpdateAnswer(&answer); err != nil {
			return echo.ErrNotFound
		}

		// user.FavoriteAnswer をインクリメント
		user := model.FindUser(&model.User{ID: answer.UID})
		user.FavoriteAnswer++
		user.FavoriteSum++
		if err := model.UpdateUser(&user); err != nil {
			return echo.ErrNotFound
		}

	} else { // いいねを取り消す
		model.DeleteAnswerGood(&model.AnswerGood{UID: uid, AID: answerID})
		answer.FavoriteCount--
		if err := model.UpdateAnswer(&answer); err != nil {
			return echo.ErrNotFound
		}

		// user.FavoriteAnswer をデクリメント
		user := model.FindUser(&model.User{ID: answer.UID})
		user.FavoriteAnswer--
		user.FavoriteSum--
		if err := model.UpdateUser(&user); err != nil {
			return echo.ErrNotFound
		}
	}

	return c.NoContent(http.StatusNoContent)
}
