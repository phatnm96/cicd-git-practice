package handlers

import (
	"cutloss-trading/app/db"
	"cutloss-trading/app/models"

	"github.com/labstack/echo/v4"
	"net/http"
)

func (h handler) GetUsers(c echo.Context) error {
	db := db.DbManager()
	users := []models.User{}
	db.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func (h handler) CreateUser(c echo.Context) error {
	db := db.DbManager()
	var err error
	user := new(models.User)
	if err = c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	result := db.Create(&user) // pass pointer of data to Create

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, user)
	}
}

func (h handler) DeleteUser(c echo.Context) error {
	db := db.DbManager()
	result := db.Delete(&models.User{}, c.Param("id"))

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, nil)
	} else {
		return c.JSON(http.StatusOK, "Delete successfully!")
	}
}
