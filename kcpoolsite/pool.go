package main

import (
	"flag"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"regexp"
	"strings"
)

var (
	addr = flag.Bool("addr", false, "find open address and print to final-port.txt")
)

// carousel data
type carousel struct {
	Name       string
	Indicators []indicator
	Items      []item
}

// slide item data
type item struct {
	Src     string
	Alt     string
	Active  string
	Caption string
}

// indicator item data
type indicator struct {
	Name   string
	Index  int
	Active string
}

type Model struct {
	C           carousel
	ScriptPaths []string
}

var templates = template.Must(template.ParseFiles("tmpl/home.tmpl", "tmpl/carousel.tmpl"))
var validPath = regexp.MustCompile("^/(home|view)/([a-zA-z0-9]+)$")
var imgPaths, err = getImgs("img/carouselImgs/")
var homeModel = buildHome()

/*********************** Init Functions *******************************/

func (c *carousel) init(imgs []string) {
	c.Name = "carousel_imgs"
	for i := range imgs {
		var active = ""
		if i == 0 {
			active = "active"
		}
		if imgs[i] == "" {
			continue
		}
		fmt.Printf("%+v\n", imgs[i])
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

func getImgs(path string) (*[]string, error) {
	var fileInfos, err = ioutil.ReadDir(path)
	pathList := make([]string, len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		if fileInfos[i].IsDir() != true && isImg(fileInfos[i].Name()) {
			pathList[i] = path + fileInfos[i].Name()
			//pathList = append(pathList, fileInfos[i].Name())
			//fmt.Printf("%+v\n", fileInfos[i].Name())
		}
	}
	return &pathList, err
}

func buildHome() *Model {
	var c carousel
	c.init(*imgPaths)
	return &Model{C: c, ScriptPaths: []string{}}
}

/************************ View Functions ******************************/

func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	/*p, err := loadModel(title)
	if err != nil {
		http.Redirect(w, r, "/home", http.StatusFound)
		return
	}
	renderTemplate(w, "view", p)*/
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "home", homeModel)
}

func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}

func renderTemplate(w http.ResponseWriter, tmpl string, m *Model) {
	err := templates.ExecuteTemplate(w, tmpl+".tmpl", m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	flag.Parse()
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/", indexHandler)
	http.Handle("/img/", http.StripPrefix("/img/", http.FileServer(http.Dir("img/"))))
	//http.Handle("/scripts/", http.StripPrefix("/scripts/", http.FileServer(http.Dir("scripts/"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("js/"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css/"))))

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
