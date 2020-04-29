package model

import "fmt"

type User struct {
	ID               int    `json:"id" gorm:"praimary_key"`
	Name             string `json:"name"`
	Password         string `json:"password"`
  TwitterId        string `json:"twitter_id"`
	FavoriteAnswer   int    `json:"favorite_answer"`
	FavoriteQuestion int    `json:"favorite_question"`
	FavoriteSum      int    `json:"favorite_sum"`

  NotificationFlag bool `json:"notification_flag"`
}

// パスワードは json として送る必要がない
type ReturnUser struct {
	ID               int    `json:"id" gorm:"praimary_key"`
	Name             string `json:"name"`
  TwitterId        string `json:"twitter_id"`
	FavoriteAnswer   int    `json:"favorite_answer"`
	FavoriteQuestion int    `json:"favorite_question"`
	FavoriteSum      int    `json:"favorite_sum"`

  NotificationFlag bool `json:"notification_flag"`
}

func (u User) IntoReturnUser() ReturnUser {
	var retUser ReturnUser
	retUser.ID = u.ID
	retUser.Name = u.Name
  retUser.TwitterId = u.TwitterId
	retUser.FavoriteAnswer = u.FavoriteAnswer
	retUser.FavoriteQuestion = u.FavoriteQuestion
	retUser.FavoriteSum = u.FavoriteSum
  retUser.NotificationFlag = u.NotificationFlag
	return retUser
}

// User の配列として定義
type Users []User
type ReturnUsers []ReturnUser

func (users Users) IntoReturnUsers() ReturnUsers {
	var retUsers ReturnUsers
	for _, s := range users {
		retUsers = append(retUsers, s.IntoReturnUser())
	}
	return retUsers
}

func CreateUser(user *User) {
	db.Create(user)
}

func FindUser(u *User) User {
	var user User
	db.Where(u).First(&user)
	return user
}

func FindUsers(u *User) Users {
	var users Users
	db.Where(u).Find(&users)
	return users
}

// 条件を満たす「質問」を, 固定長取得する
// page := ページ番号( 1-indexed ),  length := 1 ページあたりのアイテム数
func FindUsersWithPage(u *User, page int, length int, orderMode string) Users {
	var users Users
	db.Where(u).Limit(length).Offset(length * (page - 1)).Order(orderMode).Order("id desc").Find(&users)
	return users
}

// user を UPDATE
func UpdateUser(u *User) error {
	rows := db.Model(u).Update(map[string]interface{}{
		"id":                u.ID,
		"name":              u.Name,
		"password":          u.Password,
    "twitter_id":        u.TwitterId,
		"favorite_question": u.FavoriteQuestion,
		"favorite_answer":   u.FavoriteAnswer,
		"favorite_sum":      u.FavoriteSum,
    "notification_flag": u.NotificationFlag,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", u)
	}
	return nil
}

// user を 1 つ削除
func DeleteUser(u *User) error {
	if err := DeleteQuestionGood(&QuestionGood{UID: u.ID}); err != nil {
		return fmt.Errorf("質問に対するいいねを削除できませんでした")
	}

	if err := DeleteAnswerGood(&AnswerGood{UID: u.ID}); err != nil {
		return fmt.Errorf("回答に対するいいねを削除できませんでした")
	}

	if rows := db.Where(u).Delete(&User{}).RowsAffected; rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to delete", u)
	}

	return nil
}
