package handler

import (
  "fmt"
	"net/http"
	"strconv"
	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/model"
)

func PostNotification(c echo.Context) error {

  notification := new(model.Notification)
  // notification.body を bind
	if err := c.Bind(notification); err != nil {
		return err
	}

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	qid, err := strconv.Atoi(c.Param("qid"))
	if err != nil {
		return echo.ErrNotFound
	}
	t, err := strconv.Atoi(c.Param("type"))
	if err != nil {
		return echo.ErrNotFound
	}

  // uid を取得 ここでは ouid になるな(動作主なので)
  ouid := userIDFromToken(c)
  fmt.Println(ouid)
	if user := model.FindUser(&model.User{ID: ouid}); user.ID == 0 {
		return echo.ErrNotFound
	}

  notification.UID = uid
  notification.OUID = ouid
  notification.QID = qid
  notification.Type = t
  notification.Watched = false
  fmt.Println(notification)

  model.CreateNotification(notification)
  return c.JSON(http.StatusCreated, notification)
}

func GetNotification(c echo.Context) error {
	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

  items := model.FindNotifications(&model.Notification{UID: uid})
  for _, item := range items {
    item.Watched = true // 既読をつける
    model.UpdateNotification(&item)
  }

  return c.JSON(http.StatusOK, items)
}
