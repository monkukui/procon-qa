// 質問に対するいいね

package model


type AnswerGood struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"`
	AID int `json:"aid"`
}

type AnswerGoods []AnswerGood

// いいねを取得
// 例えば, UID AID 指定で，存在するかを判定できそう
func FindAnswerGoods(g *AnswerGood) AnswerGoods {
	var goods AnswerGoods
	db.Where(g).Find(&goods)
	return goods
}

// good を複数削除
func DeleteAnswerGood(g *AnswerGood) error {
  db.Where(AnswerGood{ID: g.ID}).Delete(&AnswerGood{})

	answer := FindAnswers(&Answer{ID: g.AID}, "id")[0]
  answer.FavoriteCount--
  UpdateAnswer(&answer)

  // user.FavoriteAnswer をデクリメント
  user := FindUser(&User{ID: answer.UID})
  user.FavoriteAnswer--
  user.FavoriteSum--
  UpdateUser(&user)
	return nil
}

func CreateAnswerGood(g *AnswerGood) {
	answer := FindAnswers(&Answer{ID: g.AID}, "id")[0]
  answer.FavoriteCount++
  UpdateAnswer(&answer)

  // user.FavoriteAnswer をインクリメント
  user := FindUser(&User{ID: answer.UID})
  user.FavoriteAnswer++
  user.FavoriteSum++
  UpdateUser(&user)
	db.Create(g)
}

