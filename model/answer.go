// Answer に関する定義と処理内容を記述する

package model

import "fmt"

// 回答テーブル
type Answer struct {
	UID int `json:"uid"`                    // User Id
	QID int `json:"qid"`                    // Question Id (質問に紐ずける)
	ID  int `json:"id" gorm:"praimaly_key"` // Id (インクリメント)

	// 以下, 回答の構成要素たち
	Body      string `json:"body"`
	Date      string `json:"date"`
	Favo      int    `json:"favo"`
}

// Question の配列として定義
type Answers []Answer

// 回答 を作成
func CreateAnswer(a *Answer) {
	db.Create(a)
}

// おそらくだけども, 引数はわりと自由
// 例えば, {QID: 3} を渡すと, Question に紐ずいた検索ができそう
func FindAnswers(a *Answer) Answers {
	var answers Answers
	db.Where(a).Find(&answers)
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
	if rows := db.Where(a).Delete(&Answer{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", a)
	}
	return nil
}

// answer を UPDATE
func UpdateAnswer(a *Answer) error {
	rows := db.Model(a).Update(map[string]interface{}{
		"body":      a.Body,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", a)
	}
	return nil
}
