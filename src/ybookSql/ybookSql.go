package ybookSql

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo"
	_ "github.com/lib/pq"
	"github.com/pkg/errors"
)

type Books struct {
	Isbncode	string
	Name		string
	Autor		string
	Price 		int
	Publisher	string
	Plot		string
}

var content string

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=bi18026, dbname=ybookdb password=Sshk3812-bi sslmode=disable")
	if err != nil {
		panic(err)
	}
}

func GetBooks() echo.HandlerFunc {
	return func(c echo.Context) error {
		books := Books{}
		response := ""

		rows, err := Db.Query("select isbncode, name, autor from books")
		if err != nil {
			return errors.Wrapf(err, "cannot connect SQL")
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&books.Isbncode, &books.Name, &books.Autor); err != nil {
				return errors.Wrapf(err, "cannot connect SQL")
			}
			response += "Isbncode:" + books.Isbncode + "Name:" + books.Name + "Author:" + "\n"
		}

		return c.String(http.StatusOK, response)
	}
}