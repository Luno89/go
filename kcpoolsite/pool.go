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

var templates = template.Must(template.ParseFiles("tmpl/home.html"))
var validPath = regexp.MustCompile("^/(home|view)/([a-zA-z0-9]+)$")


