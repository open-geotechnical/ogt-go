

package server


import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var SiteInfo = map[string]string{"name": "ogt-ags-go", "version": "0.1"}


func Start(address_port string) {

	// Setup www router and config modules
	router := mux.NewRouter()

	router.HandleFunc("/", H_Home)
	router.HandleFunc("/widget", H_Widget)
	router.HandleFunc("/viewer", H_Viewer)
	router.HandleFunc("/about", H_About)

	router.HandleFunc("/ajax/ags4/examples.{ext}", AX_Examples)
	router.HandleFunc("/ajax/ags4/examples", AX_Examples)


	router.HandleFunc("/ajax/ags4/parse", AX_Parse)


	router.HandleFunc("/ajax/ags4/units.{ext}", AX_Units)
	router.HandleFunc("/ajax/ags4/units", AX_Units)

	router.HandleFunc("/ajax/ags4/abbrs.{ext}",AX_Abbrs)
	router.HandleFunc("/ajax/ags4/abbrs", AX_Abbrs)

	router.HandleFunc("/ajax/ags4/abbr/{head_code}.{ext}",AX_Abbr)
	router.HandleFunc("/ajax/ags4/abbr/{head_code}", AX_Abbr)


	router.HandleFunc("/ajax/ags4/groups.{ext}",AX_Groups)
	router.HandleFunc("/ajax/ags4/groups", AX_Groups)

	router.HandleFunc("/ajax/ags4/group/{group_code}.{ext}",AX_Group)
	router.HandleFunc("/ajax/ags4/group/{group_code}", AX_Group)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}


