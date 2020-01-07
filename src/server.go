package main

import (
	"net/http"
	"html/template"
	"io"
	"io/ioutil"
	"errors"
	"encoding/json"
	"strings"
	// "fmt"		// デバッグに使ったり使わなかったり、標準出力用とか

	// 以下のライブラリはgo getする必要あり
	"github.com/labstack/echo"		// echo
	"github.com/dghubble/oauth1"	// API認証の
	"gopkg.in/mgo.v2"				// mongoDBとの接続
	// "gopkg.in/mgo.v2/bson"		// DB検索のときに使う
)

// テンプレートのレンダラーを作る
type Renderer struct {
	templates *template.Template
}

var NotFoundError = errors.New("NotFound")

func (r *Renderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return r.templates.ExecuteTemplate(w, name, data)
}

// DBから取得したデータの入る構造体
type Books struct {
	IsbnCode	string	`bson:"ISBNcode"`
	Name		string	`bson:"Name"`
	Author		string	`bson:"Author"`
	Price		string	`bson:"Price"`
	Publisher	string	`bson:"Publisher"`
	Plot		string	`bson:"Plot"`
}

func main() {
	e := echo.New()

	// "plan/*"のディレクトリを読み込む仕組み
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

	country := "23424856" // トレンドを取得する地域のIDを指定。これは日本 woeidでぐぐるとIDがいろいろ出てくる
	tags := getTrends(country) // トレンドを取得する関数を呼出

	session, _ := mgo.Dial("mongodb://localhost")	// mongoDBとの接続
		defer session.Close()
	db := session.DB("ybookDataBase")// "ybookDataBaseというデータベースを指定して接続"

	var results []Books
	db.C("Books").Find(nil).All(&results)

/*以下やり方が分からなくて全部初期化したksコード*/
	data := struct {
			Word1		string
			Word2		string
			Word3		string
			Word4		string
			Word5		string
			Word6		string
			Word7		string
			Word8		string
			Word9		string
			Word10		string
			BookName1	string
			BookName2	string
			BookName3	string
			BookName4	string
			BookName5	string
			BookName6	string
			BookName7	string
			BookName8	string
			BookName9	string
			BookName10	string
			BookName11	string
			BookName12	string
			BookName13	string
			BookName14	string
			BookName15	string
			BookName16	string
			BookName17	string
			BookName18	string
			BookName19	string
			BookName20	string
			BookName21	string
			BookName22	string
			BookName23	string
			BookName24	string
			BookName25	string
			BookName26	string
			BookName27	string
			BookName28	string
			BookName29	string
			BookName30	string

			BookPrice1	string
			BookPrice2	string
			BookPrice3	string
			BookPrice4	string
			BookPrice5	string
			BookPrice6	string
			BookPrice7	string
			BookPrice8	string
			BookPrice9	string
			BookPrice10	string
			BookPrice11	string
			BookPrice12	string
			BookPrice13	string
			BookPrice14	string
			BookPrice15	string
			BookPrice16	string
			BookPrice17	string
			BookPrice18	string
			BookPrice19	string
			BookPrice20	string
			BookPrice21	string
			BookPrice22	string
			BookPrice23	string
			BookPrice24	string
			BookPrice25	string
			BookPrice26	string
			BookPrice27	string
			BookPrice28	string
			BookPrice29	string
			BookPrice30	string

			BookPlot1	string
			BookPlot2	string
			BookPlot3	string
			BookPlot4	string
			BookPlot5	string
			BookPlot6	string
			BookPlot7	string
			BookPlot8	string
			BookPlot9	string
			BookPlot10	string
			BookPlot11	string
			BookPlot12	string
			BookPlot13	string
			BookPlot14	string
			BookPlot15	string
			BookPlot16	string
			BookPlot17	string
			BookPlot18	string
			BookPlot19	string
			BookPlot20	string
			BookPlot21	string
			BookPlot22	string
			BookPlot23	string
			BookPlot24	string
			BookPlot25	string
			BookPlot26	string
			BookPlot27	string
			BookPlot28	string
			BookPlot29	string
			BookPlot30	string
		} {
			/*tagsの中にトレンドのハッシュタグが入っているのでそれをWord1～10に代入*/
			Word1:		tags[0],
			Word2:		tags[1],
			Word3:		tags[2],
			Word4:		tags[3],
			Word5:		tags[4],
			Word6:		tags[5],
			Word7:		tags[6],
			Word8:		tags[7],
			Word9:		tags[8],
			Word10:		tags[9],

			BookName1:	results[0].Name,
			BookName2:	results[1].Name,
			BookName3:	results[2].Name,
			BookName4:	results[3].Name,
			BookName5:	results[4].Name,
			BookName6:	results[5].Name,
			BookName7:	results[6].Name,
			BookName8:	results[7].Name,
			BookName9:	results[8].Name,
			BookName10:	results[9].Name,
			BookName11:	results[10].Name,
			BookName12:	results[11].Name,
			BookName13:	results[12].Name,
			BookName14:	results[13].Name,
			BookName15:	results[14].Name,
			BookName16:	results[15].Name,
			BookName17:	results[16].Name,
			BookName18:	results[17].Name,
			BookName19:	results[18].Name,
			BookName20:	results[19].Name,
			BookName21:	results[20].Name,
			BookName22:	results[21].Name,
			BookName23:	results[22].Name,
			BookName24:	results[23].Name,
			BookName25:	results[24].Name,
			BookName26:	results[25].Name,
			BookName27:	results[26].Name,
			BookName28:	results[27].Name,
			BookName29:	results[28].Name,
			BookName30:	results[29].Name,

			BookPrice1:	results[0].Price,
			BookPrice2:	results[1].Price,
			BookPrice3:	results[2].Price,
			BookPrice4:	results[3].Price,
			BookPrice5:	results[4].Price,
			BookPrice6:	results[5].Price,
			BookPrice7:	results[6].Price,
			BookPrice8:	results[7].Price,
			BookPrice9:	results[8].Price,
			BookPrice10:	results[9].Price,
			BookPrice11:	results[10].Price,
			BookPrice12:	results[11].Price,
			BookPrice13:	results[12].Price,
			BookPrice14:	results[13].Price,
			BookPrice15:	results[14].Price,
			BookPrice16:	results[15].Price,
			BookPrice17:	results[16].Price,
			BookPrice18:	results[17].Price,
			BookPrice19:	results[18].Price,
			BookPrice20:	results[19].Price,
			BookPrice21:	results[20].Price,
			BookPrice22:	results[21].Price,
			BookPrice23:	results[22].Price,
			BookPrice24:	results[23].Price,
			BookPrice25:	results[24].Price,
			BookPrice26:	results[25].Price,
			BookPrice27:	results[26].Price,
			BookPrice28:	results[27].Price,
			BookPrice29:	results[28].Price,
			BookPrice30:	results[29].Price,

			BookPlot1:	results[0].Plot,
			BookPlot2:	results[1].Plot,
			BookPlot3:	results[2].Plot,
			BookPlot4:	results[3].Plot,
			BookPlot5:	results[4].Plot,
			BookPlot6:	results[5].Plot,
			BookPlot7:	results[6].Plot,
			BookPlot8:	results[7].Plot,
			BookPlot9:	results[8].Plot,
			BookPlot10:	results[9].Plot,
			BookPlot11:	results[10].Plot,

			BookPlot12:	results[11].Plot,
			BookPlot13:	results[12].Plot,
			BookPlot14:	results[13].Plot,
			BookPlot15:	results[14].Plot,
			BookPlot16:	results[15].Plot,
			BookPlot17:	results[16].Plot,
			BookPlot18:	results[17].Plot,
			BookPlot19:	results[18].Plot,
			BookPlot20:	results[19].Plot,
			BookPlot21:	results[20].Plot,
			BookPlot22:	results[21].Plot,
			BookPlot23:	results[22].Plot,
			BookPlot24:	results[23].Plot,
			BookPlot25:	results[24].Plot,
			BookPlot26:	results[25].Plot,
			BookPlot27:	results[26].Plot,
			BookPlot28:	results[27].Plot,
			BookPlot29:	results[28].Plot,
			BookPlot30:	results[29].Plot,
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

	e.Static("/plan2/css", "./plan2/css")
	e.Static("/plan2/pic", "./plan2/pic")

	e.Logger.Fatal(e.Start(":8080"))
}

// トレンドワードを取得する関数

func getTrends(country string) []string {
	var listHashTags []string

	/*TwitterAPIキーを入力*/

	consumerKey := "**********"
	consumerSecret := "**********"
	accessToken := "**********"
	accessSecret := "**********"

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