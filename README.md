<div align="center"><h1 align="center">Go Template Components<br /></h1>
🚨 Possibly NSFW yet.<br /><br /><br /></div>



This is a drop-in replacement for Go's `text/template` and `html/template` packages, that adds a new `{{component}}` and `{{slot}}` element to the language.

The new `component` element allows you to render reusable HTML template components that can be used across your HTML templates.

## Install

```
go get github.com/philippta/go-template@latest
```

```go
import (
    "github.com/philippta/go-template/html/template" // for html templates
    "github.com/philippta/go-template/text/template" // for text templates
)
```

## Simple example

There are two new elements to Go's templating language:
- `{{component "headline" .}} ... {{end}`
- `{{slot}}`

### Defining a component

A new component can be defined using the regular `{{define}}` element. With the help of the new `{{slot}}` element, an outlet can be set which later renders the inner/child HTML.

```html
{{define "navbar"}}
<nav class="navbar">
    <ul class="navbar-items">
        {{slot}}
    <ul>
</nav>
{{end}}
```

### Rendering a component

To render a previously defined component, the `{{component}}` element is used. The `component` element has two required parameters. First, the name of the component to render; second, the pipeline (dot).

```html
{{component "navbar" .}}
<li><a href="/">Home</a></li>
<li><a href="/about">About</a></li>
{{end}}
```

This will render the following HTML:
```html
<nav class="navbar">
    <ul class="navbar-items">
        <li><a href="/">Home</a></li>
        <li><a href="/about">About</a></li>
    <ul>
</nav>
```

### Drop-in Replacement

This package is a drop-in replacement for the `text/template` and `html/template` packages and can be used by simply replacing the imports.

```go
import (
    // "html/template"
    "github.com/philippta/go-template/html/template"
)

func main() {
    t := template.Must(template.New("layout.html").ParseFiles(
        "components/navbar.html",
        "layout.html",
    ))

    t.Execute(os.Stdout, nil)
}
```

### Passing Component Props

Passing in additional properties to a component can be achieved with a simple template function.

A full example can be found in the [example](./example) directory.

```go
func main() {
    funcs := template.FuncMap{
        "props": props,
    }

    t := template.Must(template.New("layout.html").Funcs(funcs).ParseFiles("layout.html"))
    t.Execute(os.Stdout, nil)
}

func props(v ...any) map[string]any {
	if len(v)%2 != 0 {
		panic("uneven number of key/value pairs")
	}

	m := map[string]any{}
	for i := 0; i < len(v); i += 2 {
		m[fmt.Sprint(v[i])] = v[i+1]
	}

	return m
}
```

```html
{{define "card"}}
<div class="card">
    {{if .Image}}
    <img class="card-image" src="{{.Image}}" />
    {{end}}

    <div class="card-body">
        {{slot}}
    </div>

    {{if .Link}}
    <a class="card-link" href="{{.Link}}">Open</a>
    {{end}}
</div>
{{end}}

<body>
    {{component "card" (props "Image" "/static/dog.jpeg" "Link" "/blog/dogs" )}}
    <h1>Dogs</h1>
    <p>
        This is a blog post about dogs.
    </p>
    {{end}}
</body>
```
