package main

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"goji.io/pat"
	"fmt"
	"goji.io"
	"gopkg.in/mgo.v2"
	"log"
)

type Activity struct {
	Id bson.ObjectId `json:"id" bson:"_id,omitempty"`
	City string
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func get(w http.ResponseWriter, r *http.Request)  {
	country := pat.Param(r, "country")
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	c := session.DB("travelitinerary").C("activity")
	result := Activity{}
	err = c.Find(bson.M{"country": country}).One(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("City:", result.City)

	fmt.Fprintf(w, "Corresponding City:%s", result.City)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Get("/get/:country"), get)

	http.ListenAndServe("localhost:8000", mux)
}
