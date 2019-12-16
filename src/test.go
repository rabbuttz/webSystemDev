package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)


type Books struct {
	IsbnCode	string	`bson:"ISBNcode"`
	Name		string	`bson:"Name"`
	Author		string	`bson:"Author"`
	Price		int		`bson:"Price"`
	Publisher	string	`bson:"Publisher"`
	Plot		string	`bson:"Plot"`
}

func main() {


	session, _ := mgo.Dial("mongodb://localhost")
	defer session.Clone()
	db := session.DB("ybookDataBase")

	var result []Books
	db.C("Books").Find(nil).All(&result)
	fmt.Println(result)
}

