// 通知
// 関連する質問（リンク先の質問）qid を持たせる
// type = 1：質問への回答
// type = 2：質問へのコメント
// type = 3：回答へのコメント

package model

import "fmt"

type Notification struct {
	ID  int `json:"id" gorm:"praimary_key"`
	UID int `json:"uid"` // 通知される人
	OUID int `json:"ouid"` // 関与した人
	QID int `json:"qid"`
  Type int `json:"type"`
  Watched bool `json:"watched"`
  Body string `json:"body"`
}

type Notifications []Notification

// Notification を 1 つ作成　
func CreateNotification(n *Notification) {
	db.Create(n)
}

// Notification の配列を取得
func FindNotifications(n *Notification) Notifications {
  var ret Notifications
  db.Where(n).Order("id desc").Find(&ret)
  return ret
}

// Notification を一つ削除
func DeleteNotification(n *Notification) error {
  db.Where(n).Delete(&Notification{})
  return nil
}

// notification を UPDATE
// 通知に関して，変更するのは既読フラグだけなのでこれでいい
func UpdateNotification(n *Notification) error {
	rows := db.Model(n).Update(map[string]interface{}{
    "watched": n.Watched,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", n)
	}
	return nil
}
