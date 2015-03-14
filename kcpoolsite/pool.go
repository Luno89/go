package main

import (
    "flag"
    "log"
    "html/template"
    "io/ioutil"
    "net"
    "net/http"
    "regexp"
    "errors"
)

// carousel data
type carousel struct {
	Name string
	Indicators []indicator
	Items []item
}

// slide item data
type item struct {
	Src string
	Alt string
	Active string
	Caption byte[]
}

// indicator item data
type indicator struct {
	Name string
	Index int
	Active string
}


var templates = template.Must(template.ParseFiles("tmpl/home.html","tmpl/carousel.tmpl"))
var validPath = regexp.MustCompile("^/(home|view)/([a-zA-z0-9]+)$")

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, , http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", )
}
