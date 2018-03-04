package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
	"runtime"
	"time"
)

const staticPath string = "static/"

type WebPage struct {
	Title    string
	Contents string
	// Connection *sql.DB
}

type customRouter struct {
}

func serveDynamic() {

}

func serveRendered() {

}

func serveStatic() {

}

func (customRouter) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	staticPatternString := "static/(.*)"
	templatePatternString := "template/(.*)"

	staticPattern := regexp.MustCompile(staticPatternString)
	templatePattern := regexp.MustCompile(templatePatternString)

	if staticPattern.MatchString(path) {
		serveStatic()
		page := staticPath + staticPattern.ReplaceAllString(path, "${1}") + ".html"
		http.ServeFile(rw, r, page)
	} else if templatePattern.MatchString(path) {

		serveRendered()
		urlVar := templatePattern.ReplaceAllString(path, "${1}")
		page.Contents = path
		page.Title = "This is our URL: " + urlVar
		customTemplate.Execute(rw, page)

	}

}

func gobble(s []byte) {

}

var customHTML string
var customTemplate template.Template
var page WebPage
var templateSet bool

func main() {

	runtime.GOMAXPROCS(2)

	var cr customRouter

	fileName := staticPath + "template.html"
	cH, _ := ioutil.ReadFile(fileName)
	customHTML = string(cH[:])

	page := WebPage{Title: "This is our URL: ", Contents: "Enjoy our content"}
	cT, _ := template.New("Hey").Parse(customHTML)
	customTemplate = *cT

	gobble(cH)
	log.Println(page)
	fmt.Println(customTemplate)

	server := &http.Server{
		Addr:           ":9999",
		Handler:        cr,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	server.ListenAndServe()

}
