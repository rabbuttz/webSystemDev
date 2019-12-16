package main

import (
	"net/http"
	"html/template"
	"io"
	"io/ioutil"
	"errors"
	"encoding/json"
	"strings"
	// "database/sql"
	// "fmt"

	// 以下のライブラリはgo getする必要あり
	"github.com/labstack/echo"
	"github.com/dghubble/oauth1"
	// "github.com/lib/pq"

	"ybookSql"
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

	// エラー処理（404(NotFound)と500(ServerError)のみ）
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if he, ok := err.(*echo.HTTPError); ok {
			if he.Code == 404 {
				c.Render(http.StatusNotFound,"404",nil)
			} else {
				c.Render(http.StatusInternalServerError,"500",nil)
			}
		}
	}

	country := "23424856" // トレンドを取得する地域のIDを指定。これは日本
	tags := getTrends(country) // トレンドを取得する関数を呼出

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
			/*tagsの中にトレンドのハッシュタグが入っているのでそれをWord1～10に代入*/
			Word1: tags[0],
			Word2: tags[1],
			Word3: tags[2],
			Word4: tags[3],
			Word5: tags[4],
			Word6: tags[5],
			Word7: tags[6],
			Word8: tags[7],
			Word9: tags[8],
			Word10: tags[9],
		}

	e.GET("/trend", func(c echo.Context) error {
		// plan2/trend_page.htmlがサーブされる
		// dataという構造体で単語を指定してhtmlファイルで読み込まれる
		return c.Render(http.StatusOK, "trend_page", data)
	})

	e.GET("/cart", func(c echo.Context) error {
		return c.Render(http.StatusOK, "cart", nil)
	})

	e.GET("/trend/rank1", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank1", data)
	})
	e.GET("/trend/rank2", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank2", data)
	})
	e.GET("/trend/rank3", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank3", data)
	})
	e.GET("/trend/rank4", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank4", data)
	})
	e.GET("/trend/rank5", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank5", data)
	})
	e.GET("/trend/rank6", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank6", data)
	})
	e.GET("/trend/rank7", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank7", data)
	})
	e.GET("/trend/rank8", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank8", data)
	})
	e.GET("/trend/rank9", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank9", data)
	})
	e.GET("/trend/rank10", func(c echo.Context) error {
		return c.Render(http.StatusOK, "rank10", data)
	})
	e.GET("/sql", ybookSql.GetBooks())

	// 静的ファイルのため
	e.Static("/plan2/css", "./plan2/css")
	e.Static("/plan2/pic", "./plan2/pic")

	e.Logger.Fatal(e.Start(":8080"))
}

// トレンドのタグを取得する関数
func getTrends(country string) []string {
	var listHashTags []string

	/*APIキーを入力*/
	consumerKey := "339ttP5v0cWVKlBmkIeyfSjmk"
	consumerSecret := "MZJXV2SbTiNxeD0yf7dhortxX69IdP78ADlP4kPxrN9jdFssiV"
	accessToken := "955371524403298304-sKHG03JGLUoIJ13m9LgwtLWyacT3K4f"
	accessSecret := "BjtC6S7ger7RiOm8gGjo0MbUUn6UuYUajOvmnleCUK0FN"

	// APIキーが未入力時のエラー出力
	if consumerKey == "" || consumerSecret == "" || accessToken == "" || accessSecret == "" {
		panic("Missing required environment variable")
	}

	config := oauth1.NewConfig(consumerKey, consumerSecret)
	token := oauth1.NewToken(accessToken, accessSecret)

	httpClient := config.Client(oauth1.NoContext, token)

	// トレンド取得の地域指定
	path := "https://api.twitter.com/1.1/trends/place.json?id=" + country

	resp, _ := httpClient.Get(path)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	// JSON形式の[]をカット
	var JSON = strings.TrimLeft(string(body), "[")
	JSON = strings.TrimRight(JSON, "]")

	var info map[string]interface{}
	json.Unmarshal([]byte(JSON), &info)
	trends := info["trends"].([]interface{})

	for _, element := range trends {
		if trendList, ok := element.(map[string]interface{}); ok {
			for key, value := range trendList {
				if strings.Contains(key, "name") && strings.Contains(value.(string), "") {
					listHashTags = append(listHashTags, value.(string))
				}
			}
		}
	}
	return listHashTags
}

/*
func getSql() (bk []books) {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5432 user=bi18026 password=Sshk3812-bi dbname=ybookdb sslmode=disable")
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}

	rows, err := db.Query("SELECT * FROM books")
	if err != nil {
		fmt.Println(err)
	}

	var bk []books
	for rows.Next() {
		var b books
		rows.Scan(&b.name, &b.plot)
		bk = append(bk, b)
	}
	return bk
}
*/


/*
func viewHashTagHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/trend" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		getPage, _ := template.ParseFiles("trend_page.html")
		getPage.Execute(w, r)
	case "POST":
		var country string
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		country = r.FormValue("country")
		tags := getHashTags(country)
		postPage, _ := template.ParseFiles("trend_page.html")
		postPage.Execute(w, tags)
	default:
		fmt.Fprintf(w, "Unable to get result.")
	}
}
*/