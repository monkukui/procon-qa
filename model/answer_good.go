// 質問に対するいいね

package model
import "fmt"

type AnswerGood struct {
  ID       int `json:"id" gorm:"praimaly_key"`
  UID      int `json:"uid" gorm:"praimaly_key"`
  AID      int `json:"aid" gorm:"praimaly_key"`
}

type AnswerGoods []AnswerGood

func CreateAnswerGood(g * AnswerGood) {
  db.Create(g)
}

// いいねを取得
// 例えば, UID AID 指定で，存在するかを判定できそう
func FindAnswerGoods(g *AnswerGood) AnswerGoods {
	var goods AnswerGoods
	db.Where(g).Find(&goods)
	return goods
}

// good を 1 つ削除
func DeleteAnswerGood(g *AnswerGood) error {
	if rows := db.Where(g).Delete(&AnswerGood{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", g)
	}
	return nil
}
