package server

import (
	//"fmt"
	"encoding/json"
	"net/http"
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

// SendAjaxPayload is the function that sends the http reply
// machine encoded `payload` formatted,  ie a serialiser
// html should not hit here, but otherwise expected is
// a reply with the "bites" in the particular machine readable format
// and correct mime type etc eg json, yaml and xml.hell in m$.excel
func SendAjaxPayload(resp http.ResponseWriter, request *http.Request, payload interface{}) {

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
			bites, err = json.MarshalIndent(payload, "", "  ")
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

type ErrorPayload struct {
	Success bool   ` json:"success" `
	Error   string ` json:"error" `
}

func SendAjaxError(resp http.ResponseWriter, request *http.Request, err error) {
	SendAjaxPayload(resp, request, ErrorPayload{Success: true, Error: err.Error()})
}

type UnitsPayload struct {
	Success bool        ` json:"success" `
	Units   []ags4.Unit ` json:"units" `
}

// handles /ags/4/units.*
func AX_Units(resp http.ResponseWriter, req *http.Request) {

	payload := new(UnitsPayload)
	payload.Success = true
	payload.Units = ags4.Units

	SendAjaxPayload(resp, req, payload)
}

var EndPoints = map[string]string{
	"/":            "Data and Sys information",
	"/ags/4/all":   "AGS4: All data",
	"/ags/4/units": "AGS4: Units",
}

func AX_Info(resp http.ResponseWriter, req *http.Request) {
	payload := map[string]interface{}{
		"repos":      "https://bitbucket.org/daf0dil/ags-def-json",
		"version":    "0.1-alpha",
		"server_utc": time.Now().UTC().Format("2006-01-02 15:04:05"),
		"endpoints":  EndPoints,
	}
	SendAjaxPayload(resp, req, payload)
}

type AbbrevsPayload struct {
	Success bool           ` json:"success" `
	Abbrevs []*ags4.Abbrev ` json:"abbreviations" `
}

// handles /ags/4/groups.
func AX_Abbrevs(resp http.ResponseWriter, req *http.Request) {

	var e error
	payload := new(AbbrevsPayload)
	payload.Success = true
	payload.Abbrevs, e = ags4.GetAbbrevs()
	if e != nil {
		SendAjaxError(resp, req, e)
		return
	}
	SendAjaxPayload(resp, req, payload)
}

type GroupsPayload struct {
	Success bool          ` json:"success" `
	Groups  []*ags4.Group ` json:"groups" `
}

// handles /ags/4/groups.
func AX_Groups(resp http.ResponseWriter, req *http.Request) {

	var e error
	payload := new(GroupsPayload)
	payload.Success = true
	payload.Groups, e = ags4.GetGroups()
	if e != nil {
		SendAjaxError(resp, req, e)
		return
	}
	SendAjaxPayload(resp, req, payload)
}

type AbbrevPayload struct {
	Success bool         ` json:"success" `
	Found   bool         ` json:"found" `
	Abbrev  *ags4.Abbrev ` json:"abbreviation" `
}

// handles /ags/4/units.*
func AX_Abbrev(resp http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	abbr, found, err := ags4.GetAbbrev(vars["head_code"])
	if err != nil {
		SendAjaxError(resp, req, err)
		return
	}
	payload := new(AbbrevPayload)
	payload.Success = true
	payload.Found = found
	payload.Abbrev = abbr
	SendAjaxPayload(resp, req, payload)
}

type GroupPayload struct {
	Success bool        ` json:"success" `
	Group   *ags4.Group ` json:"group" `
}

// handles /ags/4/units.*
func AX_Group(resp http.ResponseWriter, req *http.Request) {

	vars := mux.Vars(req)
	grp, err := ags4.GetGroup(vars["group_code"])
	if err != nil {
		SendAjaxError(resp, req, err)
		return
	}
	payload := new(GroupPayload)
	payload.Success = true
	payload.Group = grp
	SendAjaxPayload(resp, req, payload)
}
