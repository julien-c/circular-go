package main

import (
	"encoding/json"
	"labix.org/v2/mgo"
	"net/http"
)

type Counter struct {
	Count int `json:"count"`
}

var session, _ = mgo.Dial("localhost")
var c = session.DB("circular").C("posts")

func serve(w http.ResponseWriter, r *http.Request) {
	count, _ := c.Count()
	m := Counter{count}
	b, _ := json.Marshal(m)

	w.Header().Set("Content-Type", "application/json")
	w.Write(b)
}

func main() {
	http.HandleFunc("/api/counter", serve)
	http.ListenAndServe(":8080", nil)
}
