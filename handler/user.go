package handler

import (
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)


func GetUser(c echo.Context) error {

	UserID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

	user := model.FindUser(&model.User{ID: UserID})

	return c.JSON(http.StatusOK, user)
}
