
package server

import (
	"net/http"

	"github.com/flosch/pongo2"

)

var tplHome = pongo2.Must( pongo2.FromFile("templates/home.html") )

// Home page url's =  / /home /index.html
func H_Home(resp http.ResponseWriter, request *http.Request){

	err := tplHome.ExecuteWriter(pongo2.Context{"query": request.FormValue("query")}, resp)
	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
	}
}
