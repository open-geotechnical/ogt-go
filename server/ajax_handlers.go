

package server


import (
	//"fmt"
	"net/http"
	"encoding/json"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"bitbucket.org/daffodil/ags2go/ags4"

)

/* WTF ???? said pedro
SupportedFormats  := map[string]string{
	"json":	"application/json",
	"js":	"application/json",
	"yaml":	"text/yaml",
	"txt":	"text/plain",
	"ags4":	"text/plain???",
}
*/


// SendPayload is the "main function" that sends the http reply
// machine encoded `payload` formatted,  ie a serialiser
// html should not hit here, but otherwise expected is
// a reply with the "bites" in the particular machine readable format
// and correct mime type etc eg json, yaml and xml.hell in m$.excel
func SendPayload(resp http.ResponseWriter, request *http.Request, payload interface{}  ) {

	// pretty returns indents data and readable (notably json) is ?pretty=1 in url
	pretty := request.URL.Query().Get("pretty") == "1"

	// Determine which encoding from the mux/router
	vars := mux.Vars(request)
	enc := vars["ext"]
	// TODO validate encoding and serialiser
	// eg yaml, json/js, html, xlsx, ags4,

	// TODO map[string] = encoding


	// Lets get ready to encode folks...
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

	} else {
		bites = []byte("OOPes No `.ext` set ")
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

func H_Units(resp http.ResponseWriter, req *http.Request){

	payload := new(UnitsPayload)
	payload.Success = true
	payload.Units = ags4.Units

	SendPayload(resp, req,  payload)
}

