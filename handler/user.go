package handler

import (
	"net/http"
	"strconv"
  //"fmt"
	"github.com/labstack/echo"
	// c-color さんに依存しちゃってる...
	// なんとか再現性があるように作り変えることができないもんかね...?
	// でも, wifi なくても起動できるんだよなぁ
	// どこを参照しているんだろう.....
	// "github.com/x-color/simple-webapp/model"
	// ? できたのか ?
	"github.com/monkukui/procon-qa/model"
)


func GetUser(c echo.Context) error {

	uid := userIDFromToken(c)
	if user := model.FindUser(&model.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	UserID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

	user := model.FindUser(&model.User{ID: UserID})

	return c.JSON(http.StatusOK, user)
}
