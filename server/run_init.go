

package server


import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)




func Start(address_port string) {

	// Setup www router and config modules
	router := mux.NewRouter()

	router.HandleFunc("/", AX_Info)
	//router.HandleFunc("/info.{ext}", AX_Info)

	router.HandleFunc("/ags/4/units.{ext}", AX_Units)
	router.HandleFunc("/ags/4/units", AX_Units)

	//router.HandleFunc("/ags/4/units.{ext}", H_Units)


	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}


