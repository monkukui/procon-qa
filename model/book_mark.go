// 質問に対するブックマーク

package model

import "fmt"

type BookMark struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
	QID int `json:"qid"`
}

type BookMarks []BookMark

func CreateBookMark(b *BookMark) {
	db.Create(b)
}

// いいねを取得
// 例えば, UID QID 指定で，存在するかを判定できそう
func FindBookMarks(b *BookMark) BookMarks {
  var marks BookMarks
	db.Where(b).Order("id desc").Find(&marks)
	return marks
}

// good を 1 つ削除
func DeleteBookMark(b *BookMark) error {
	if rows := db.Where(b).Delete(&BookMark{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", b)
	}
	return nil
}
