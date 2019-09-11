// Question に関する定義と処理内容を記述する

package model

import "fmt"

type Question struct {
  UID       int    `json:"uid"`
  TID       int    `json:"tid"`
	ID        int    `json:"id" gorm:"praimaly_key"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	Url       string `json:"url"`
  Date      string `json:"date"`
	Completed bool   `json:"completed"`
}

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
  fmt.Println(questions)
  return questions
}

// questions を 1 つ削除
func DeleteQuestion(q *Question) error {
	if rows := db.Where(q).Delete(&Question{}).RowsAffected; rows == 0 {
    return fmt.Errorf("Could not find Todo (%v) to delete", q)
  }
  return nil
}

// question を UPDATE
func UpdateQuestion(q *Question) error {
  rows := db.Model(q).Update(map[string]interface{}{
    "title": q.Title,
    "body": q.Body,
    "url": q.Url,
    "completed": q.Completed,
  }).RowsAffected
  if rows == 0 {
    return fmt.Errorf("Could not find Todo (%v) to update", q)
  }
  return nil
}
