package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

var session *mgo.Session

type ArqHtml struct {
	Id   bson.ObjectId `bson:"_id" json:"id"`
	Body string        `json:"body"`
}

func main() {
	session, _ = mgo.Dial("mongodb://mega:megamega@kahana.mongohq.com:10093/MegaMetricas")
	collection := session.DB("MegaMetricas").C("PlSql")
	var arqhtml ArqHtml
	arqhtml.Id = bson.NewObjectId()
	var a, _ = ioutil.ReadFile("C:/temp/teste.svg")
	arqhtml.Body = string(a)
	fmt.Println(arqhtml.Id)
	err := collection.Insert(arqhtml)
	fmt.Println(err)
	var html ArqHtml
	err = collection.FindId(arqhtml.Id).One(&html)
	fmt.Println(err)
}
