package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/monkukui/procon-qa/server/entity"
)

// users のサイズを取得する
func GetUserSize(c echo.Context) error {
	return c.JSON(http.StatusOK, len(entity.FindUsers(&entity.User{})))
}

func GetUser(c echo.Context) error {
	UserID, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}
	user := entity.FindUser(&entity.User{ID: UserID})
	return c.JSON(http.StatusOK, user.IntoReturnUser())
}

// ユーザをページ取得する
func GetUsersWithPage(c echo.Context) error {

	PageID, err := strconv.Atoi(c.Param("page")) // ページ番号 (1-indexed)
	modeId, err := strconv.Atoi(c.Param("mode")) // ソートの設定
	PageLength := 20                             // 1 ページあたりの長さ

	if err != nil {
		return echo.ErrNotFound
	}
	mode := "favorite_sum desc"

	if modeId == 2 {
		mode = "favorite_question desc"
	} else if modeId == 3 {
		mode = "favorite_answer desc"
	}

	users := entity.FindUsersWithPage(&entity.User{}, PageID, PageLength, mode)
	return c.JSON(http.StatusOK, users.IntoReturnUsers())
}

// user を削除する
func DeleteUser(c echo.Context) error {
	uid := userIDFromToken(c)
	if user := entity.FindUser(&entity.User{ID: uid}); user.ID == 0 {
		return echo.ErrNotFound
	}

	userId, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

	if uid != userId {
		fmt.Println("他人のアカウントです")
	}

	if err := entity.DeleteUser(&entity.User{ID: uid}); err != nil {
		fmt.Println("削除できませんでした")
	}

	return c.NoContent(http.StatusNoContent)
}

func UpdateUser(c echo.Context) error {
	uid := userIDFromToken(c)
	user := entity.FindUser(&entity.User{ID: uid})
	if user.ID == 0 {
		return echo.ErrNotFound
	}

	postUser := new(entity.User)
	if err := c.Bind(postUser); err != nil {
		return err
	}
	if postUser.Name == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name",
		}
	}
	if u := entity.FindUser(&entity.User{Name: postUser.Name}); u.ID != 0 && user.Name != u.Name {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

	user.Name = postUser.Name
	user.TwitterId = postUser.TwitterId

	fmt.Println(user)

	entity.UpdateUser(&user)
	return c.JSON(http.StatusOK, user.IntoReturnUser())
}

func UpdateUserNotification(c echo.Context) error {

	uid, err := strconv.Atoi(c.Param("uid"))
	if err != nil {
		return echo.ErrNotFound
	}

	flag, err := strconv.Atoi(c.Param("flag"))
	if err != nil {
		return echo.ErrNotFound
	}
	user := entity.FindUser(&entity.User{ID: uid})
	user.NotificationFlag = (flag == 1)
	entity.UpdateUser(&user)
	return c.JSON(http.StatusOK, user.IntoReturnUser())
}
