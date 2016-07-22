

package server


import (
	//"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"bitbucket.org/daffodil/ags2go/ags4"

)


// sends the  `payload` formatted, json/yaml types for now
func SendAjaxPayload(resp http.ResponseWriter, request *http.Request, payload interface{}  ) {

	// return indented data (notably json) is ?pretty=1 in url
	pretty := request.URL.Query().Get("pretty") == "1"

	// Determine which encoding from the mux/router
	vars := mux.Vars(request)
	enc := vars["ext"]

	var bites []byte
	var err error
	var mime string = "text/plain"

	if enc == "yaml" {
		bites, err = yaml.Marshal(payload)
		mime = "text/yaml"

	} else if enc == "json" || enc == "js" {
		if pretty {
			bites, err = json.MarshalIndent(payload, "", "    ")
		} else {
			bites, err = json.Marshal(payload)
		}
		mime = "application/json"
	}

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", mime)
	resp.Write(bites)
}


type UnitsPayload struct {
	Success bool 		` json:"success" `
	Units []ags4.Unit 	` json:"units" `
}

func H_AjaxUnits(resp http.ResponseWriter, req *http.Request){

	payload := new(UnitsPayload)
	payload.Success = true
	payload.Units = ags4.Units

	SendAjaxPayload(resp, req,  payload)
}
