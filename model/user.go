package model


type User struct {
	ID       int    `json:"id" gorm:"praimaly_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

// User の配列として定義
type Users []User

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