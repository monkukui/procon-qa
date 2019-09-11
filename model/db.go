package model

import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB

func init() {
    var err error
    db, err = gorm.Open("sqlite3", "db/sample.db")
    if err != nil {
      panic("failed to connect database")
    }
    db.AutoMigrate(&User{})
    db.AutoMigrate(&Todo{})
    db.AutoMigrate(&Question{}) // Question データベースを作成してみる
}
