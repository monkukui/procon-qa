package model
import "fmt"

type User struct {
	ID       int    `json:"id" gorm:"praimaly_key"`
	Name     string `json:"name"`
	Password string `json:"password"`
	FavoriteAnswer int    `json:"favoriteAnswer"`
	FavoriteQuestion int  `json:"favoriteQuestion"`
	FavoriteSum int `json:"favoriteSum"`
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


// user を UPDATE
func UpdateUser(u *User) error {
	rows := db.Model(u).Update(map[string]interface{}{
		"id":        u.ID,
		"name":      u.Name,
		"password":  u.Password,
		"favoriteQuestion":  u.FavoriteQuestion,
		"favoriteAnswer":    u.FavoriteAnswer,
		"favoriteSum": u.FavoriteSum,
	}).RowsAffected
	if rows == 0 {
		return fmt.Errorf("Could not find Todo (%v) to update", u)
	}
	return nil
}
