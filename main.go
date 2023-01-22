package main

import (
	"mongo-golang/controller"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc, err := controller.NewUserController(getSession())
	if err != nil {
		panic(err)
	}
	r.GET("/user/:userID", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:userID", uc.DeleteUser)
	http.ListenAndServe("localhost:9000", r)
}

func getSession() *mgo.Session {
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
