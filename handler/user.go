package handler

import (
	"net/http"
	"strconv"
	// "fmt"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

// users のサイズを取得する
func GetUserSize(c echo.Context) error {
  return c.JSON(http.StatusOK, len(model.FindUsers(&model.User{})))
}

func GetUser(c echo.Context) error {
	UserID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	user := model.FindUser(&model.User{ID: UserID})
	return c.JSON(http.StatusOK, user)
}

// ユーザをページ取得する
func GetUsersWithPage(c echo.Context) error {

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	modeId, err := strconv.Atoi(c.Param("mode"))        // ソートの設定
	PageLength := 10                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}
	mode := "favorite_sum desc"

    if modeId == 2 {
      mode = "favorite_question desc"
    } else if modeId == 3 {
      mode = "favorite_answer desc"
    }
	users := model.FindUsersWithPage(&model.User{}, PageID, PageLength, mode)
	return c.JSON(http.StatusOK, users)
}
