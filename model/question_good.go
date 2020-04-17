// 質問に対するいいね

package model

import "fmt"

type QuestionGood struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
	QID int `json:"qid"`
}

type QuestionGoods []QuestionGood

// いいねを取得
// 例えば, UID QID 指定で，存在するかを判定できそう
func FindQuestionGoods(g *QuestionGood) QuestionGoods {
	var goods QuestionGoods
	db.Where(g).Find(&goods)
	return goods
}

// good を 1 つ削除
func DeleteQuestionGood(g *QuestionGood) error {
	if rows := db.Where(g).Delete(&QuestionGood{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", g)
	}

	question := FindQuestions(&Question{ID: g.QID})[0]
  question.FavoriteCount--
  UpdateQuestion(&question)

  // user.FavoriteQuestion をインクリメント
  user := FindUser(&User{ID: question.UID})
  user.FavoriteQuestion--
  user.FavoriteSum--
  UpdateUser(&user)
	return nil
}

// good を 1 つ作成　
func CreateQuestionGood(g *QuestionGood) {
	question := FindQuestions(&Question{ID: g.QID})[0]

  question.FavoriteCount++
  UpdateQuestion(&question)

  // user.FavoriteQuestion をインクリメント
  user := FindUser(&User{ID: question.UID})
  user.FavoriteQuestion++
  user.FavoriteSum++
  UpdateUser(&user)
	db.Create(g)
}
