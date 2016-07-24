

package server


import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)




func Start(address_port string) {

	// Setup www router and config modules
	router := mux.NewRouter()

	router.HandleFunc("/", H_Home)
	router.HandleFunc("/info.{ext}", H_Info)

	router.HandleFunc("/ags/4/units.{ext}", A_Units)
	router.HandleFunc("/ags/4/units", A_Units)

	//router.HandleFunc("/ags/4/units.{ext}", H_Units)


	fmt.Println("Serving on " + address_port)
	http.Handle("/", router)
	http.ListenAndServe(address_port , nil)
}



func H_Info(resp http.ResponseWriter, req *http.Request){
	payload := map[string]interface{} {
		"pong": "yipee",
		"ts":   "--timestamp_here--",
		"client_ip": "client_ip",
		"ags2go": "Versoin 0.16 ",
		"a_num": 20.1356223,
		"a_init": 3256,
	}
	SendPayload(resp, req, payload)
}
