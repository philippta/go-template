package main

import (
	"log"
	"os"

	"github.com/philippta/go-template/text/template"
)

func main() {
	tmpl, err := template.New("layout.html").ParseFiles(
		"components/navbar.html",
		"layout.html",
		"home.html",
	)

	template.Must(tmpl, err)
	if err := tmpl.Execute(os.Stdout, nil); err != nil {
		log.Fatal(err)
	}
}
