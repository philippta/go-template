package main

import (
	"embed"
	"fmt"
	"log"
	"os"

	"github.com/philippta/go-template/html/template"
)

//go:embed templates
var FS embed.FS

func main() {
	funcs := template.FuncMap{
		"props": props,
	}

	tmpl := template.Must(template.New("home.html").Funcs(funcs).ParseFS(FS,
		"templates/components/card.html",
		"templates/components/navbar.html",
		"templates/home.html",
	))

	if err := tmpl.Execute(os.Stdout, "<script>alert()</script>"); err != nil {
		log.Fatal(err)
	}
}

func props(v ...any) map[string]any {
	if len(v)%2 != 0 {
		panic("uneven number of key/value pairs")
	}

	m := map[string]any{}
	for i := 0; i < len(v); i += 2 {
		key := fmt.Sprint(v[i])
		val := v[i+1]
		m[key] = val
	}

	return m
}
