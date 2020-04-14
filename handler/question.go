package handler

import (
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
	"net/http"
	"strconv"
	"time"
  "fmt"
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
	if err != nil {
		return echo.ErrNotFound
	}
	modeId, err := strconv.Atoi(c.Param("mode")) // ソートの設定
	if err != nil {
		return echo.ErrNotFound
	}

	PageLength := 10 // 1 ページあたりの長さ
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

// pageId, userId で質問をページ取得する
func GetUserQuestions(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

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

func GetUserQuestionSize(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}
  fmt.Println(len(model.FindQuestions(&model.Question{UID: uid})))
  return c.JSON(http.StatusOK, len(model.FindQuestions(&model.Question{UID: uid})))
}

// pageId, userId で質問をページ取得する
func GetUserAnswers(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	PageLength := 10                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}

  // ユーザーの回答を取得
  answers := model.FindAnswers(&model.Answer{UID: uid}, "id")

  // このユーザが関与した質問の qid リストを重複無しで構築
  var qidList []int

  for _, answer := range answers {
    qidList = append(qidList, answer.QID)
  }

  // map を使って重複削除
  mp := make(map[int]bool)
  var uniqQidList []int

  for _, id := range qidList {
    if !mp[id] {
      mp[id] = true
      uniqQidList = append(uniqQidList, id)
    }
  }


  // PageLength と PageID を使ってよしなに範囲を決定する
  lb := (PageID - 1) * PageLength
  if lb >= len(uniqQidList) {
    return echo.ErrNotFound
  }

  var questions model.Questions
  for i := 0; i < PageLength && lb + i < len(uniqQidList); i++ {
	  q := model.FindQuestions(&model.Question{ID: uniqQidList[lb + i]})
    if (len(q) != 1) {
      return echo.ErrNotFound
    }
    questions = append(questions, q[0])
  }

	return c.JSON(http.StatusOK, questions)
}

func GetUserAnswerSize(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

  // ユーザーの回答を取得
  answers := model.FindAnswers(&model.Answer{UID: uid}, "id")

  // このユーザが関与した質問の qid リストを重複無しで構築
  var qidList []int

  for _, answer := range answers {
    qidList = append(qidList, answer.QID)
  }

  // map を使って重複削除
  mp := make(map[int]bool)
  userAnswerSize := 0

  for _, id := range qidList {
    if !mp[id] {
      mp[id] = true
      userAnswerSize++
    }
  }

  return c.JSON(http.StatusOK, userAnswerSize)
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

// いいねをする（or 取り消す）
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

	// question_good の FindQuestionGoods を呼び出して，いいねがされているかを取得する

	// いいねがされていたら
	// DeleteQuestionGood
	// question と user の favoritecount をデクリメント

	// いいねがされていなかったら，
	// CreateQuestionGood
	// question と user の favoritecount をインクリメント
	goods := model.FindQuestionGoods(&model.QuestionGood{UID: uid, QID: questionID})

	if len(goods) == 0 { // いいねをする
		model.CreateQuestionGood(&model.QuestionGood{UID: uid, QID: questionID})

		question.FavoriteCount++
		if err := model.UpdateQuestion(&question); err != nil {
			return echo.ErrNotFound
		}
		// user.FavoriteQuestion をインクリメント
		user := model.FindUser(&model.User{ID: question.UID})
		user.FavoriteQuestion++
		user.FavoriteSum++
		if err := model.UpdateUser(&user); err != nil {
			return echo.ErrNotFound
		}

	} else { // いいねを取り消す
		model.DeleteQuestionGood(&model.QuestionGood{UID: uid, QID: questionID})
		question.FavoriteCount--
		if err := model.UpdateQuestion(&question); err != nil {
			return echo.ErrNotFound
		}
		// user.FavoriteQuestion をインクリメント
		user := model.FindUser(&model.User{ID: question.UID})
		user.FavoriteQuestion--
		user.FavoriteSum--
		if err := model.UpdateUser(&user); err != nil {
			return echo.ErrNotFound
		}
	}

	return c.NoContent(http.StatusNoContent)
}

// ブックマークをする（or 取り消す）
func BookMarkQuestion(c echo.Context) error {
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

	marks := model.FindBookMarks(&model.BookMark{UID: uid, QID: questionID})

	if len(marks) == 0 { // ブックマークをする
		model.CreateBookMark(&model.BookMark{UID: uid, QID: questionID})
	} else { // いいねを取り消す
		model.DeleteBookMark(&model.BookMark{UID: uid, QID: questionID})
	}

	return c.NoContent(http.StatusNoContent)
}

// ブックマークされた質問を取得する
func GetBookMarkedQuestions(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid")) // ページ番号 (1-indexed)
	if err != nil {
		return echo.ErrNotFound
	}

  books := model.FindBookMarks(&model.BookMark{UID: uid})

  var questions model.Questions
  for _, b := range books {
	  q := model.FindQuestions(&model.Question{ID: b.QID})
    if (len(q) != 1) {
      return echo.ErrNotFound
    }
    questions = append(questions, q[0])
  }

	return c.JSON(http.StatusOK, questions)
}

// 閲覧数をインクリメント
func BrowseQuestion(c echo.Context) error {
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

	// URL チェック
	// javascript::alert(1) みたいなのを弾く
	// https:// か http:// のみにする
	// 文字列長による面倒な分岐を避けるために，ダミー文字を 10 こ付け加える
	checkUrl := question.Url + "xxxxxxxxxx"

	if checkUrl != "xxxxxxxxxx" && checkUrl[0:7] != "http://" && checkUrl[0:8] != "https://" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid url fields",
		}
	}
	// if question.Url[]

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	question.UID = uid
	question.Completed = false
	question.AnswerCount = 0
	question.FavoriteCount = 0
	question.BrowseCount = 0

	question.Date = time.Now().Format("2006/01/02 15:04:05")

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
		return echo.ErrNotFound
	}

	// 先に関連する answer を全て削除
	model.DeleteAnswer(&model.Answer{QID: questionID})

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
