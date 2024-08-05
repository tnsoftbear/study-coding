package main

import (
	"html/template"
	"io"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func NewTemplates() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Count struct {
	Count int
}

func main() {
	e := echo.New()
	e.Renderer = NewTemplates()
	e.Use(middleware.Logger())

	counter := Count{Count: 0}

	e.GET("/", func(c echo.Context) error {
		return c.Render(200, "index", counter)
	})

	e.POST("/count", func(c echo.Context) error {
		counter.Count++
		return c.Render(200, "count", counter)
	})

	e.Logger.Fatal(e.Start(":42069"))
}
