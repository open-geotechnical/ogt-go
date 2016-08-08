
package server

import (
	"net/http"

	"github.com/flosch/pongo2"

)

type NavItem struct {
	Url string
	Title string
}

var Nav []NavItem

func init(){

	addNav("/", "Home")
	addNav("/about", "About")
	addNav("/widget", "Widget")
	addNav("/view", "View")
	addNav("/ags4", "AGS4")
}
func addNav(url, title string)  {
	Nav = append(Nav, NavItem{Url: url, Title: title})
}


func NewContext() pongo2.Context {
	c := pongo2.Context{"BOOTSTRAP": "bootstrap-3.3.7", "EXT_VER": "ext-4.2.1-gpl",
						"site": SiteInfo, "LOAD_EXT": false}
	return c
}

func RenderTemplate(resp http.ResponseWriter, request *http.Request, tpl *pongo2.Template, ctx pongo2.Context) {

	err := tpl.ExecuteWriter(ctx, resp)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
}

var tplHome = pongo2.Must( pongo2.FromFile("templates/home.html") )
var tplAbout = pongo2.Must( pongo2.FromFile("templates/about.html") )
var tplWidget = pongo2.Must( pongo2.FromFile("templates/widget.html") )
var tplView = pongo2.Must( pongo2.FromFile("templates/view.html") )

// Home page
func H_Home(resp http.ResponseWriter, request *http.Request){

	c := NewContext()
	RenderTemplate(resp, request, tplHome, c)
}

// About page TODO
func H_About(resp http.ResponseWriter, request *http.Request){
	c := NewContext()
	RenderTemplate(resp, request, tplAbout, c)
}

// Widget page
func H_Widget(resp http.ResponseWriter, request *http.Request){
	c := NewContext()
	c["LOAD_EXT"] = true
	RenderTemplate(resp, request, tplWidget, c)
}

// View AGS page
func H_View(resp http.ResponseWriter, request *http.Request){
	c := NewContext()
	c["LOAD_EXT"] = true
	RenderTemplate(resp, request, tplView, c)
}
