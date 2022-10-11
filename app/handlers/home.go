package handlers

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func (h handler) HomePage(c echo.Context) error {
	return c.Render(http.StatusOK, "home.html", map[string]interface{}{
		"name": "Cutloss Trading",
		"msg":  "something cool",
	})
}
