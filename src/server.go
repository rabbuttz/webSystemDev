package main

import (
	"net/http"
	"html/template"
	"io"
	"errors"

	"github.com/labstack/echo"
)

// テンプレートのレンダラーを作る
type Renderer struct {
	templates *template.Template
}

var NotFoundError = errors.New("NotFound")

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// plan/*を読み込む仕組み
	e.Renderer = &Renderer{
		templates: template.Must(template.ParseGlob("plan2/*.html")),
	}

	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound,"404",nil)
			} else {
				c.Render(http.StatusInternalServerError,"500",nil)
			}
		}
	}

	e.GET("/trend", func(c echo.Context) error {
		// plan2/trend_page.htmlがサーブされる
		return c.Render(http.StatusOK, "trend_page", nil)
	})

	e.GET("/cart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "ContentInTheCart", nil)
	})

	// 静的ファイルのため
	e.Static("/plan2/css", "./plan2/css")
	e.Static("/plan2/pic", "./plan2/pic")

	e.Logger.Fatal(e.Start(":8080"))
}