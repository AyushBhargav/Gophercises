package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

// Story which is unmarshalled from the JSON.
type Story struct {
	Title   string   `json:"title"`
	Story   []string `json:"story"`
	Options []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}

// Adventure key value pair for title and story
type Adventure map[string]Story

// NewHandler for http requests
func NewHandler(adventure Adventure) http.Handler {
	return handler{adventure}
}

type handler struct {
	adventure Adventure
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	dataHTML, err := ioutil.ReadFile("adventure.htm")
	if err != nil {
		panic(err)
	}

	t := template.Must(template.New("").Parse(string(dataHTML)))
	path := r.URL.Path[1:]
	if len(path) == 0 {
		path = "intro"
	}
	err = t.Execute(w, h.adventure[path])
	if err != nil {
		panic(err)
	}
}

func main() {
	data, err := ioutil.ReadFile("story.json")
	if err != nil {
		panic("Can't read story.json")
	}

	var adventure Adventure
	err = json.Unmarshal(data, &adventure)

	if err != nil {
		panic(err)
	}

	port := flag.Int("port", 3000, "Port for start adventure application")
	handler := NewHandler(adventure)

	fmt.Printf("Starting app at Port: %d\n", *port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), handler))
}
