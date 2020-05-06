package model

import "fmt"

type QuestionComment struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
  QID int `json:"qid"`

	Body  string `json:"body"`
	Date  string `json:"date"`
}

type QuestionComments []QuestionComment

// 作成
func CreateQuestionComment(q *QuestionComment) {
  db.Create(q)
}

// 複数取得
func FindQuestionComments(q *QuestionComment) QuestionComments {
  var comments QuestionComments
  db.Where(q).Find(&comments)
  return comments
}

// 一つ削除
func DeleteQuestionComment(q *QuestionComment) error {

  // 関連して消すべきものが今のところはない
  db.Where(QuestionComment{ID: q.ID}).Delete(&Question{})
  return nil
}

// 更新
func UpdateQuestionComment(q *QuestionComment) error {
  rows := db.Model(q).Update(map[string]interface{}{
    "body":       q.Body,
  }).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", q)
	}
  return nil
}
