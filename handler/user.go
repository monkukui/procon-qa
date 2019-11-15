package handler

import (
	"net/http"
	"strconv"
	// "fmt"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

// answers のサイズを取得する
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
