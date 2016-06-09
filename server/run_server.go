

package server


import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/daffodil/ags2go"
)


func Start() {

	// Setup www router and config modules
	router := mux.NewRouter()

	router.HandleFunc("/ajax/ags/", AGSAjaxHandler)

	addr := "0.0.0.0:1558"
	fmt.Println("Serving on " + addr)
	http.Handle("/", router)
	http.ListenAndServe(addr , nil)
}

