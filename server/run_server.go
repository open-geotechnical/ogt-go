

package server


import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	//"bitbucket.org/daffodil/ags2go"
	//"bitbucket.org/daffodil/ags2go/ags4"
)

func PingTestHandler(http.ResponseWriter, *http.Request){

}


func Start(address_port string) {

	// Setup www router and config modules
	router := mux.NewRouter()

	router.HandleFunc("/ajax/ping/", PingTestHandler)
	router.HandleFunc("/ags/4/units.{ext}", H_AjaxUnits)


	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}
