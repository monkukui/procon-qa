package entity

import "fmt"

type AnswerComment struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
	AID int `json:"aid"`

	Body string `json:"body"`
	Date string `json:"date"`
}

type AnswerComments []AnswerComment

// 作成
func CreateAnswerComment(a *AnswerComment) {
	db.Create(a)
}

// 複数取得
func FindAnswerComments(a *AnswerComment) AnswerComments {
	var comments AnswerComments
	db.Where(a).Find(&comments)
	return comments
}

// 一つ削除
func DeleteAnswerComment(a *AnswerComment) error {

	// 関連して消すべきものが今のところはない
	db.Where(AnswerComment{ID: a.ID}).Delete(&AnswerComment{})
	return nil
}

// 更新
func UpdateAnswerComment(a *AnswerComment) error {
	rows := db.Model(a).Update(map[string]interface{}{
		"body": a.Body,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", a)
	}
	return nil
}
