// Question に関する定義と処理内容を記述する

package model

import "fmt"

// 質問テーブル
type Question struct {
	UID int `json:"uid"`                    // User Id
	TID int `json:"tid"`                    // Tag Id
	ID  int `json:"id" gorm:"praimary_key"` // Id (インクリメント)

	// 以下, 質問の構成要素たち
	Title string `json:"title"`
	Body  string `json:"body"`
	Url   string `json:"url"`
	Date  string `json:"date"`

	Completed     bool `json:"completed"`
	AnswerCount   int  `json:"answerCount"`
	FavoriteCount int  `json:"favoriteCount"`
	BrowseCount   int  `json:"browseCount"`
}

// Question の配列として定義
type Questions []Question

// question を作成
func CreateQuestion(q *Question) {
	db.Create(q)
}

// おそらくだけども, 引数はわりと自由
// 例えば, {UID: 3} を渡すと, 絞り込みで取得ができるっぽい
func FindQuestions(q *Question) Questions {
	var questions Questions
	db.Where(q).Find(&questions)
	return questions
}

// 条件を満たす「質問」を, 固定長取得する
// page := ページ番号( 1-indexed ),  length := 1 ページあたりのアイテム数
func FindQuestionsWithPage(q *Question, page int, length int, orderMode string) Questions {
	var questions Questions
	db.Where(q).Limit(length).Offset(length * (page - 1)).Order(orderMode).Order("id desc").Find(&questions)
	return questions
}

// question を 1 つ削除
func DeleteQuestion(q *Question) error {

	// 関連する answer を全て削除
	answers := FindAnswers(&Answer{QID: q.ID}, "id")
	for _, answer := range answers {
		// DeleteAnswer(&answer) これだとダメ id だけ渡すようにしなきゃ
		DeleteAnswer(&Answer{ID: answer.ID, QID: answer.QID})
	}

	// 関連する good を全て削除
	goods := FindQuestionGoods(&QuestionGood{QID: q.ID})
	for _, good := range goods {
		DeleteQuestionGood(&QuestionGood{ID: good.ID, QID: q.ID})
	}

	// 関連するコメントを全て削除
	comments := FindQuestionComments(&QuestionComment{QID: q.ID})
	for _, comment := range comments {
		DeleteQuestionComment(&QuestionComment{ID: comment.ID})
	}

	// question を削除
	db.Where(Question{ID: q.ID}).Delete(&Question{})
	return nil
}

// question を UPDATE
func UpdateQuestion(q *Question) error {
	rows := db.Model(q).Update(map[string]interface{}{
		"title":         q.Title,
		"body":          q.Body,
		"url":           q.Url,
		"date":          q.Date,
		"completed":     q.Completed,
		"answerCount":   q.AnswerCount,
		"favoriteCount": q.FavoriteCount,
		"browseCount":   q.BrowseCount,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", q)
	}
	return nil
}
