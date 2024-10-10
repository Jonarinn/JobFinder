package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

func main() {

	var rootHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		log.Printf("Running on %s", r.URL)

		content := parsePage("root")
		header := parsePage("header")

		data := struct {
			Header  template.HTML
			Content template.HTML
		}{
			Header:  template.HTML(header),
			Content: template.HTML(content),
		}

		tmpl.Execute(w, data)
	}

	var jobHandler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hi"))
	}

	http.HandleFunc("/jobs", jobHandler)
	http.HandleFunc("/", rootHandler)

	http.ListenAndServe(":8000", nil)
}

func parsePage(name string) template.HTML {
	page, err := os.ReadFile(fmt.Sprintf("pages/%s/index.html", name))
	if err != nil {
		log.Fatalf("error parsing %s")
	}
	return template.HTML(page)
}
