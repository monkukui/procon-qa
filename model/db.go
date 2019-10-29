package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
    // _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
  // db, err = gorm.Open("sqlite3", "db/hoge.db")
  db, err = gorm.Open("postgres", "user=postgres password=postgres dbname=huga sslmode=disable")
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Question{}) // Question データベースを作成してみる
	db.AutoMigrate(&Answer{}) // Answer データベースを作成してみる
}
