// 質問に対するいいね

package model

import "fmt"

type QuestionGood struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
	QID int `json:"qid"`
}

type QuestionGoods []QuestionGood

func CreateQuestionGood(g *QuestionGood) {
	db.Create(g)
}

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
	return nil
}
