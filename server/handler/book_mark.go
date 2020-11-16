package handler

import (
	"net/http"
	"strconv"

	// "fmt"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/server/entity"
)

func QuestionBookMarked(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		return echo.ErrNotFound
	}

	goods := entity.FindBookMarks(&entity.BookMark{UID: uid, QID: qid})
	return c.JSON(http.StatusOK, len(goods) != 0)
}
