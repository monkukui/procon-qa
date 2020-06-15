package handler

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

func QuestionFavorited(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		return echo.ErrNotFound
	}

	goods := model.FindQuestionGoods(&model.QuestionGood{UID: uid, QID: qid})
	return c.JSON(http.StatusOK, len(goods) != 0)
}

func CountQuestionGood(c echo.Context) error {
	return c.JSON(http.StatusOK, len(model.FindQuestionGoods(&model.QuestionGood{})))
}
