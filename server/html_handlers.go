
package server

import (
	"net/http"

	"github.com/flosch/pongo2"

)

func NewContext() pongo2.Context {
	c := pongo2.Context{"BOOTSTRAP": "bootstrap-3.3.7", "EXT_VER": "ext-4.2.1-gpl",
						"site": SiteInfo}
	return c
}

func RenderTemplate(resp http.ResponseWriter, request *http.Request, tpl *pongo2.Template, ctx pongo2.Context) {

	err := tpl.ExecuteWriter(ctx, resp)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
}

var tplHome = pongo2.Must( pongo2.FromFile("templates/home.html") )
var tplWidget = pongo2.Must( pongo2.FromFile("templates/widget.html") )

// Home page
func H_Home(resp http.ResponseWriter, request *http.Request){

	c := NewContext()
	RenderTemplate(resp, request, tplHome, c)
}

// Widget page
func H_Widget(resp http.ResponseWriter, request *http.Request){

	c := NewContext()
	RenderTemplate(resp, request, tplWidget, c)
}