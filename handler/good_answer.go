package handler

import (
	"net/http"
	"strconv"
	// "fmt"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

func AnswerFavorited(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	aid, err := strconv.Atoi(c.Param("aid"))
	if err != nil {
		return echo.ErrNotFound
	}

	goods := model.FindAnswerGoods(&model.AnswerGood{UID: uid, AID: aid})
	return c.JSON(http.StatusOK, len(goods) != 0)
}

func CountAnswerGood(c echo.Context) error {
  return c.JSON(http.StatusOK, len(model.FindAnswerGoods(&model.AnswerGood{})))
}
