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
		data := struct {
			Word1 string
			Word2 string
			Word3 string
			Word4 string
			Word5 string
			Word6 string
			Word7 string
			Word8 string
			Word9 string
			Word10 string
		} {
			Word1: "A",
			Word2: "B",
			Word3: "C",
			Word4: "D",
			Word5: "E",
			Word6: "F",
			Word7: "G",
			Word8: "H",
			Word9: "I",
			Word10: "J",
		}
		return c.Render(http.StatusOK, "trend_page", data)
	})

	e.GET("/cart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "cart", nil)
	})

	e.GET("/trend/rank1", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank1", nil)
	})
	e.GET("/trend/rank2", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank2", nil)
	})
	e.GET("/trend/rank3", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank3", nil)
	})
	e.GET("/trend/rank4", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank4", nil)
	})
	e.GET("/trend/rank5", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank5", nil)
	})
	e.GET("/trend/rank6", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank6", nil)
	})
	e.GET("/trend/rank7", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank7", nil)
	})
	e.GET("/trend/rank8", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank8", nil)
	})
	e.GET("/trend/rank9", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank9", nil)
	})
	e.GET("/trend/rank10", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank10", nil)
	})

	// 静的ファイルのため
	e.Static("/plan2/css", "./plan2/css")
	e.Static("/plan2/pic", "./plan2/pic")

	e.Logger.Fatal(e.Start(":8080"))
}