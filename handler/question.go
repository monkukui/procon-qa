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
  return c.JSON(http.StatusOK, len(model.FindQuestions(&model.Question{})))
}

// 解決済みの質問のサイズを取得
func GetCompletedQuestionSize(c echo.Context) error {
  return c.JSON(http.StatusOK, len(model.FindQuestions(&model.Question{Completed: true})))
}

// 質問をページ取得する
func GetQuestionsWithPage(c echo.Context) error {
	// 質問の取得にユーザの情報はいらない
  // uid := userIDFromToken(c)
	// if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
	// 	return echo.ErrNotFound
	// }

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	modeId, err := strconv.Atoi(c.Param("mode"))        // ソートの設定
	PageLength := 10                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}
	mode := "id desc"

    if modeId == 2 {
      mode = "answer_count desc"
    } else if modeId == 3 {
      mode = "browse_count desc"
    } else if modeId == 4 {
      mode = "favorite_count desc"
    } else if modeId == 5 {
      mode = "completed"
    }

    fmt.Println(mode)
	questions := model.FindQuestionsWithPage(&model.Question{}, PageID, PageLength, mode)
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

  questions := model.FindQuestionsWithPage(&model.Question{UID: uid}, PageID, PageLength, "id")
	return c.JSON(http.StatusOK, questions)
}

// 質問を 1 つ 取得する
func GetQuestion(c echo.Context) error {
  // 認証必要なし
	// uid := userIDFromToken(c)
	// if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
	// 	return echo.ErrNotFound
	// }

	QuestionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

	question := model.FindQuestions(&model.Question{ID: QuestionID})[0]
	return c.JSON(http.StatusOK, question)
}

// いいねをする
func FavoriteQuestion(c echo.Context) error {
  uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

  questions := model.FindQuestions(&model.Question{ID: questionID})
  if len(questions) == 0 {
    return echo.ErrNotFound
  }

  question := questions[0]
  if question.UID == uid {
    // 自分の質問にいいねはできません
    return echo.ErrNotFound
  }
  question.FavoriteCount++
  if err := model.UpdateQuestion(&question); err != nil {
    return echo.ErrNotFound
  }
  return c.NoContent(http.StatusNoContent)
}

// 閲覧数をインクリメント
func BrowseQuestion(c echo.Context) error {
    fmt.Println("in go")
  questionID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}

  questions := model.FindQuestions(&model.Question{ID: questionID})
  if len(questions) == 0 {
    return echo.ErrNotFound
  }

  question := questions[0]
  question.BrowseCount++
  if err := model.UpdateQuestion(&question); err != nil {
    return echo.ErrNotFound
  }
  return c.NoContent(http.StatusNoContent)
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

    // uid が question の uid と一致していなければダメ
    if uid != model.FindQuestions(&model.Question{ID: questionID})[0].UID {
        fmt.Println("他人の質問です")
        return echo.ErrNotFound
    }

	// 先に関連する answer を全て削除
    model.DeleteAnswer(&model.Answer{QID: questionID});

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