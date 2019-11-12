package handler

import (
	"net/http"
	"strconv"
  "fmt"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

// 質問を全取得する
func GetAllQuestions(c echo.Context) error {
	// todos := model.FindTodos(&model.Todo{UID: uid})
	// &model.Question{} とすることで, 条件なしで取得する <=> 全取得 となる
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
	questions := model.FindQuestions(&model.Question{})
	return c.JSON(http.StatusOK, questions)
}

// questions のサイズを取得する
func GetQuestionSize(c echo.Context) error {
  questions := model.FindQuestions(&model.Question{})
  cnt := len(questions)
  return c.JSON(http.StatusOK, cnt)
}

// 質問をページ取得する
func GetQuestionsWithPage(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	PageLength := 10                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}

	questions := model.FindQuestionsWithPage(&model.Question{}, PageID, PageLength)
	return c.JSON(http.StatusOK, questions)
}

// pageId で質問をページ取得する
func GetUserQuestionsWithPage(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	PageLength := 10                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}

  questions := model.FindQuestionsWithPage(&model.Question{UID: uid}, PageID, PageLength)
	return c.JSON(http.StatusOK, questions)
}

// 質問を 1 つ 取得する
func GetQuestion(c echo.Context) error {

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	QuestionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	question := model.FindQuestions(&model.Question{ID: QuestionID})[0]
	return c.JSON(http.StatusOK, question)
}

// 質問を投稿する
func PostQuestion(c echo.Context) error {
	question := new(model.Question)

	// question に 送信されてきたデータを bind している
	if err := c.Bind(question); err != nil {
		return err
	}

	// 妥当性判定
	// Title, Body が空欄ではないことをチェックする
	if question.Title == "" || question.Body == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid to or message fields",
		}
	}

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	question.UID = uid
  question.Completed = false;
	model.CreateQuestion(question)

	return c.JSON(http.StatusCreated, question)
}

// 質問を削除する
func DeleteQuestion(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	// ID: questionID, UID: uid とすることで, 別のユーザが他人の投稿を削除できないようになってる
	if err := model.DeleteQuestion(&model.Question{ID: questionID, UID: uid}); err != nil {
		return echo.ErrNotFound
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateQuestionCompleted(c echo.Context) error {

  uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

  questions := model.FindQuestions(&model.Question{ID: questionID, UID: uid})
  fmt.Println(len(questions))
  if len(questions) == 0 {
    return echo.ErrNotFound
  }
  question := questions[0]
  question.Completed = !questions[0].Completed
  if err := model.UpdateQuestion(&question); err != nil {
    return echo.ErrNotFound
  }

  return c.NoContent(http.StatusNoContent)
}
