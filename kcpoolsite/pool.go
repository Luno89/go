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
var imgPaths = getImgs("data/carousel_imgs")

func (c *carousel) init(imgs *[]string) {
	c.Name = "carousel_imgs"
	for i := 0; i < len(imgs); i++ {
		var active = ""
		if i == 0 {
			active = "active"
		}
		var temp := item{imgs[i], "image", active, ""}
		var indicator := indicator{c.Name, i, active}
		append(c.Items, temp)
		append(c.Indicators, indicator)
	}
}

func isImg(s string) (bool) {
	switch s {
		case Contains(s, "jpg"):
			return true
		case Contains(s, "png"):
			return true
		case Conatins(s, "tif"):
			return true
		case Contains(s, "gif"):
			return true
	}
	return false
}

func getImgs(path string) (*[]string, error){
	var fileInfos = ReadDir(path)
	var pathList := make([]string, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		if !fileInfos[i].IsDir && isImg(fileInfos[i].Name) {
			pathList = append(pathList, fileInfos[i].Name)
		}
	}
	return &pathList
}

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
