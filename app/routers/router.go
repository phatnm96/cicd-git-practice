package routers

import (
	"gorm.io/gorm"

	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"cutloss-trading/app/handlers"
)

// Define the template registry struct
type TemplateRegistry struct {
	templates *template.Template
}

// Implement e.Renderer interface
func (t *TemplateRegistry) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func Init(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	handler := handlers.New(db)

	// Instantiate a template registry and register all html files inside the view folder
	e.Renderer = &TemplateRegistry{
		templates: template.Must(template.ParseGlob("app/public/views/*.html")),
	}

	e.GET("/", handler.HomePage)
	e.GET("/users", handler.GetUsers)
	e.POST("/users", handler.CreateUser)
	e.DELETE("/users/:id", handler.DeleteUser)

	return e
}
