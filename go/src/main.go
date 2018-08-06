package main

import (
	"fmt"
	"goji.io"
	"goji.io/pat"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
)

type Activity struct {
	Id                   bson.ObjectId `json:"id" bson:"_id,omitempty"`
	Duration_Hours       string
	Rating               string
	Activity_Level       string
	Summary              string
	Detailed_Description string
	Province             string
	City                 string
	Country              string
}

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func get(w http.ResponseWriter, r *http.Request) {
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
	fmt.Println("City:", result.Id)
	fmt.Println("Duration:", result.Duration_Hours)
	fmt.Println("Rating:", result.Rating)
	fmt.Println("Activity Level:", result.Activity_Level)
	fmt.Println("Summary:", result.Summary)
	fmt.Println("Detailed Description:", result.Detailed_Description)
	fmt.Println("Province:", result.Province)
	fmt.Println("City:", result.City)
	fmt.Println("Country:", result.Country)

	fmt.Fprintf(w, "Id:%s", result.Id)
	fmt.Fprintf(w, "Duration:%s", result.Duration_Hours)
	fmt.Fprintf(w, "Rating:%s", result.Rating)
	fmt.Fprintf(w, "Activity Level:%s", result.Activity_Level)
	fmt.Fprintf(w, "Summary:%s", result.Summary)
	fmt.Fprintf(w, "Detailed Description:%s", result.Detailed_Description)
	fmt.Fprintf(w, "Province:%s", result.Province)
	fmt.Fprintf(w, "City:%s", result.City)
	fmt.Fprintf(w, "Country:%s", result.Country)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Get("/get/:country"), get)

	http.ListenAndServe("localhost:8000", mux)
}
