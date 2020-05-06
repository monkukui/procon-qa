// Answer に関する定義と処理内容を記述する

package model

import "fmt"


// 回答テーブル
type Answer struct {
	ID  int `json:"id" gorm:"praimary_key"` // Id (インクリメント)
	UID int `json:"uid"`                    // User Id
	QID int `json:"qid"`                    // Question Id (質問に紐ずける)

	// 以下, 回答の構成要素たち
	Body          string `json:"body"`
	Date          string `json:"date"`
	FavoriteCount int    `json:"favoriteCount"`
}

// Question の配列として定義
type Answers []Answer

// 回答 を作成
func CreateAnswer(a *Answer) {
	questions := FindQuestions(&Question{ID: a.QID})
	// questions[0] の 回答数をインクリメント
	questions[0].AnswerCount++
	UpdateQuestion(&questions[0])
	db.Create(a)
}

// おそらくだけども, 引数はわりと自由
// 例えば, {QID: 3} を渡すと, Question に紐ずいた検索ができそう
func FindAnswers(a *Answer, orderMode string) Answers {
	var answers Answers
	db.Where(a).Order(orderMode).Order("id desc").Find(&answers)
	return answers
}

// 条件を満たす「質問」を, 固定長取得する
// page := ページ番号( 1-indexed ),  length := 1 ページあたりのアイテム数
func FindAnswersWithPage(a *Answer, page int, length int) Answers {
	var answers Answers
	db.Where(a).Limit(length).Offset(length * (page - 1)).Find(&answers)
	return answers
}

// answer を 1 つ削除
func DeleteAnswer(a *Answer) error {

  // 関連する good を削除
  goods := FindAnswerGoods(&AnswerGood{AID: a.ID})
  for _, good := range goods {
    // DeleteAnswerGood(&good) ダメ，id 指定
    DeleteAnswerGood(&AnswerGood{ID: good.ID})
  }

  // 関連するコメントを全て削除
  comments := FindAnswerComments(&AnswerComment{AID: a.ID})
  for _, comment := range comments {
    // DeleteAnswerComment(&comment) ダメ，id 指定してくだいさい
    DeleteAnswerComment(&AnswerComment{ID: comment.ID})
  }

  // answer を削除
	db.Where(a).Delete(&Answer{})

	questions := FindQuestions(&Question{ID: a.QID})
	// questions[0] の 回答数をデクリメント
	questions[0].AnswerCount--
	UpdateQuestion(&questions[0])
	return nil
}

// answer を UPDATE
func UpdateAnswer(a *Answer) error {
	rows := db.Model(a).Update(map[string]interface{}{
		"body":          a.Body,
		"favoriteCount": a.FavoriteCount,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", a)
	}
	return nil
}
