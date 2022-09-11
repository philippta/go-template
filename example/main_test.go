package main_test

import (
	"log"
	"os"
	"testing"

	"github.com/philippta/go-template/text/template"
)

func TestXxx(t *testing.T) {
	tmpl, err := template.New("layout.html").ParseFiles(
		"components/navbar.html",
		"layout.html",
	)

	template.Must(tmpl, err)
	if err := tmpl.Execute(os.Stdout, "pipe"); err != nil {
		log.Fatal(err)
	}
}
