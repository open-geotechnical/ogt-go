
package server

import (
	"net/http"

	"github.com/flosch/pongo2"

)

type NavItem struct {
	Url string
	Title string
	//Selected bool
}
type Page struct {
	Title string
	Url string
}

var Nav []NavItem

func init(){
	addNav("/", "Home")
	addNav("/about", "About")
	addNav("/ags4", "AGS4 Data Dict")
	addNav("/viewer", "Viewer")
	//addNav("/ags4", "AGS4")
}
func addNav(url, title string)  {
	Nav = append(Nav, NavItem{Url: url, Title: title})
}


func NewContext(url string) pongo2.Context {

	page := Page{Title: "==NO TITLE==", Url: "#"}
	for _, nav := range Nav {
		if nav.Url == url {
			page.Url = nav.Url
			page.Title = nav.Title
			break;
		}
	}

	c := pongo2.Context{	"BOOTSTRAP": "bootstrap-3.3.7",
							"EXT_VER": "ext-4.2.1-gpl",
							"LOAD_EXT": false,
							"site": SiteInfo,
							"nav": Nav,
							"page": page}
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
var tplView = pongo2.Must( pongo2.FromFile("templates/viewer.html") )

// Home page
func H_Home(resp http.ResponseWriter, request *http.Request){

	c := NewContext("/")
	RenderTemplate(resp, request, tplHome, c)
}

// About page TODO
func H_About(resp http.ResponseWriter, request *http.Request){
	c := NewContext("/about")
	RenderTemplate(resp, request, tplAbout, c)
}

// Widget page
func H_Widget(resp http.ResponseWriter, request *http.Request){
	c := NewContext("/widget")
	c["LOAD_EXT"] = true
	RenderTemplate(resp, request, tplWidget, c)
}

// View AGS page
func H_Viewer(resp http.ResponseWriter, request *http.Request){
	c := NewContext("/viewer")
	c["LOAD_EXT"] = true
	RenderTemplate(resp, request, tplView, c)
}
