package model

import (
    "os"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres" // Use PostgreSQL in gorm
    "github.com/lib/pq"
    // _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
	var err error
  // db, err = gorm.Open("sqlite3", "db/hoge.db")

  url := os.Getenv("DATABASE_URL")
  env := ""
  stage := ""
  stage = os.Getenv("STAGE")
  if "" == url {
    env = "development"
  } else {
    env = "production"
  }
  if stage == "TEST" {
    env = "test"
  }
  if env == "development" {
    db, err = gorm.Open("postgres", "user=postgres password=postgres dbname=huga sslmode=disable")
  } else if env == "production" {

    connection, err := pq.ParseURL(url)
    if err != nil {
      panic(err.Error())
    }
    connection += " sslmode=require"
    db, err = gorm.Open("postgres", connection);
  } else {
    db, err = gorm.Open("postgres", "host=postgres user=postgres password=postgres dbname=postgres sslmode=disable")
  }
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Question{}) // Question データベースを作成してみる
	db.AutoMigrate(&Answer{}) // Answer データベースを作成してみる
	db.AutoMigrate(&QuestionGood{}) // Answer データベースを作成してみる
	db.AutoMigrate(&AnswerGood{}) // Answer データベースを作成してみる
}
