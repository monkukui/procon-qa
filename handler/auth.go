package handler

import (
	"net/http"
	"time"

  "golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/monkukui/procon-qa/model"
	// "github.com/x-color/simple-webapp/model"
)

type jwtCustomClaims struct {
	UID  int    `json:"uid"`
	Name string `json:"name"`
	jwt.StandardClaims
}

var signingKey = []byte("secret")

var Config = middleware.JWTConfig{
	Claims:     &jwtCustomClaims{},
	SigningKey: signingKey,
}

func Signup(c echo.Context) error {
	user := new(model.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	if user.Name == "" || user.Password == "" {
		return &echo.HTTPError{
			Code:    http.StatusBadRequest,
			Message: "invalid name or password",
		}
	}

	if u := model.FindUser(&model.User{Name: user.Name}); u.ID != 0 {
		return &echo.HTTPError{
			Code:    http.StatusConflict,
			Message: "name already exists",
		}
	}

  hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
  if err != nil {
    return err
  }
  user.Password = string(hash)
	model.CreateUser(user)
	user.Password = ""
	return c.JSON(http.StatusCreated, user)
}

func Login(c echo.Context) error {
	u := new(model.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	user := model.FindUser(&model.User{Name: u.Name})
	if user.ID == 0 {
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid name",
		}
	}
  if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(u.Password)); err != nil {
    // 失敗
		return &echo.HTTPError{
			Code:    http.StatusUnauthorized,
			Message: "invalid password",
		}
  }

	claims := &jwtCustomClaims{
		user.ID,
		user.Name,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString(signingKey)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, map[string]string{
		"token": t,
	})
}

func userIDFromToken(c echo.Context) int {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwtCustomClaims)
	uid := claims.UID
	return uid
}
