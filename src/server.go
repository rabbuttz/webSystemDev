package main

import (
	"net/http"
	"github.com/labstack/echo"
	"html/template"
	"io"
)

// テンプレートのレンダラーを作る
type Renderer struct {
	templates *template.Template
}

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()

	// plan/*を読み込む仕組み
	e.Renderer = &Renderer{
		templates: template.Must(template.ParseGlob("plan2/*.html")),
	}

	e.GET("/trend", func(c echo.Context) error {
		// plan2/trend_page.htmlがサーブされる
		return c.Render(http.StatusOK, "trend_page", nil)
	})

	// 静的ファイルの:w
	ため
	e.Static("/plan2/css", "./plan2/css")
	e.Static("/plan2/pic", "./plan2/pic")

	e.Logger.Fatal(e.Start(":8080"))
}