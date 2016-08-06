

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

	router.HandleFunc("/ags/4/abbreviations.{ext}",AX_Abbrevs)
	router.HandleFunc("/ags/4/abbreviations", AX_Abbrevs)
	router.HandleFunc("/ags/4/abbrevs.{ext}",AX_Abbrevs)
	router.HandleFunc("/ags/4/abbrevs", AX_Abbrevs)

	router.HandleFunc("/ags/4/abbrev/{head_code}.{ext}",AX_Abbrev)
	router.HandleFunc("/ags/4/abbrev/{head_code}", AX_Abbrev)


	router.HandleFunc("/ags/4/groups.{ext}",AX_Groups)
	router.HandleFunc("/ags/4/groups", AX_Groups)

	router.HandleFunc("/ags/4/group/{group_code}.{ext}",AX_Group)
	router.HandleFunc("/ags/4/group/{group_code}", AX_Group)


	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}


