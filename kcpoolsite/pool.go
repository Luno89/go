package main

import (
    "flag"
    "log"
    "html/template"
    "io/ioutil"
    "net"
    "net/http"
    "regexp"
    "strings"
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
	Caption string
}

// indicator item data
type indicator struct {
	Name string
	Index int
	Active string
}

type page struct {
	
}

var templates = template.Must(template.ParseFiles("tmpl/home.html","tmpl/carousel.tmpl"))
var validPath = regexp.MustCompile("^/(home|view)/([a-zA-z0-9]+)$")
var imgPaths, err = getImgs("img/carouselImgs/")

func (c *carousel) init(imgs []string) {
	c.Name = "carousel_imgs"
	for i, value := range imgs {
		var active = ""
		if i == 0 {
			active = "active"
		}
		temp := item{imgs[i], "image", active, ""}
		indicator := indicator{c.Name, i, active}
		c.Items = append(c.Items, temp)
		c.Indicators = append(c.Indicators, indicator)
	}
}

func isImg(imgName string) bool {
	var s = strings.ToLower(imgName)
	switch {
		case strings.Contains(s, "jpg"):
			return true
		case strings.Contains(s, "png"):
			return true
		case strings.Contains(s, "tif"):
			return true
		case strings.Contains(s, "gif"):
			return true
	}
	return false
}

func getImgs(path string) (*[]string, error){
	var fileInfos, err = ioutil.ReadDir(path)
	pathList := make([]string, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		if fileInfos[i].IsDir() != true && isImg(fileInfos[i].Name()) {
			pathList = append(pathList, fileInfos[i].Name())
		}
	}
	return &pathList, err
}

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", "")
}

func makeHandler(fn func (http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w,r)
			return
		}
		fn(w, r, m[2])
	}		
}

func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/", indexHandler)
	
	if *addr {
		l, err := net.Listen("tcp", "127.0.0.1:0")
		if err != nil {
			log.Fatal(err)
		}
		err = ioutil.WriteFile("final-port.txt", []byte(l.Addr().String()), 644)
		if err != nil {
			log.Fatal(err)
		}
		s := &http.Server{}
		s.Serve(l)
		return
	}
	http.ListenAndServe(":8080", nil)
}
