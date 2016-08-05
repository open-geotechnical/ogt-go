

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

	router.HandleFunc("/ags/4/groups.{ext}",AX_Groups)
	router.HandleFunc("/ags/4/groups", AX_Groups)

	router.HandleFunc("/ags/4/group/{code}.{ext}",AX_Group)
	router.HandleFunc("/ags/4/group/{code}", AX_Group)


	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}


