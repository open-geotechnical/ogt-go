

package server


import (
	//"fmt"
	"net/http"
	"encoding/json"
	"time"

	"github.com/gorilla/mux"
	"gopkg.in/yaml.v2"

	"bitbucket.org/daf0dil/ags2go/ags4"


)

/* WTF ???? said pedro
SupportedFormatsMaps  := map[string]string{
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
	pretty := true //request.URL.Query().Get("pretty") == "0"

	// Determine which encoding from the mux/router
	vars := mux.Vars(request)
	enc := vars["ext"]
	if enc == "" {
		enc = "json"
	}
	// TODO validate encoding and serialiser
	// eg yaml, json/js, html, xlsx, ags4,

	// TODO map[string] = encoding


	// Lets get ready to encode folks...
	var bites []byte
	var err error
	var content_type string = "text/plain"

	if enc == "yaml" {
		bites, err = yaml.Marshal(payload)
		content_type = "text/yaml"

	} else if enc == "json" || enc == "js" {
		if pretty {
			bites, err = json.MarshalIndent(payload, "", "    ")
		} else {
			bites, err = json.Marshal(payload)
		}
		content_type = "application/json"

	} else {
		bites = []byte("OOPs no `.ext` recognised ")
	}

	if err != nil {
		http.Error(resp, err.Error(), http.StatusInternalServerError)
		return
	}

	resp.Header().Set("Content-Type", content_type)
	resp.Write(bites)
}


type UnitsPayload struct {
	Success bool 		` json:"success" `
	Units []ags4.Unit 	` json:"units" `
}

// handles /ags/4/units.*
func AX_Units(resp http.ResponseWriter, req *http.Request){

	payload := new(UnitsPayload)
	payload.Success = true
	payload.Units = ags4.Units

	SendPayload(resp, req,  payload)
}

var EndPoints = map[string]string {
	"/": "Data and Sys information",
	"/ags/4/all": "AGS4: All data",
	"/ags/4/units": "AGS4: Units",
}


func AX_Info(resp http.ResponseWriter, req *http.Request){
	payload := map[string]interface{} {
			"repos": "https://bitbucket.org/daf0dil/ags-def-json",
			"version": "0.1-alpha",
			"server_utc":   time.Now().UTC().Format("2006-01-02 15:04:05"),
			"endpoints": EndPoints,

	}
	SendPayload(resp, req, payload)
}
