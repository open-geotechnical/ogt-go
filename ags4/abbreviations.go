package ags4

import (
	"encoding/json"
	//"fmt"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)

// this is a lookup for autocomplete etc
var abbrsMap map[string]*Abbr

func init() {
	abbrsMap = make(map[string]*Abbr)
}

type AbbrevItem struct {
	Item        string ` json:"item" `
	Description string ` json:"description" `
	DateAdded   string ` json:"date_added" `
	AddedBy     string ` json:"added_by" `
	Status      string ` json:"status" `
}

// An AGS abbreviation
type Abbr struct {
	HeadCode    string       ` json:"head_code" `
	Description string       ` json:"description" `
	Items       []AbbrevItem ` json:"items" `
}

func GetAbbrevs() ([]*Abbr, error) {

	var keys []string
	for k := range abbrsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	abbrevs := make([]*Abbr, 0, 0)
	for _, k := range keys {
		//fmt.Println("Key:", k, "Value:", m[k])
		abbrevs = append(abbrevs, abbrsMap[k])
	}
	return abbrevs, nil
}

func GetAbbrev(heading_code string) (*Abbr, bool, error) {
	heading_code = strings.TrimSpace(heading_code)
	if len(heading_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	ab, found := abbrsMap[heading_code]
	return ab, found, nil
}



type abbrevReader struct {
	Info  abbrevInfoReader   ` json:"info" `
	Items []abbrevItemReader ` json:"meta" `
}

type abbrevInfoReader struct {
	Group     string ` json:"ABBR_HDNG" `
	GroupDesc string ` json:"group" `
	Heading   string ` json:"heading" `
}
type abbrevItemReader struct {
	Item        string ` json:"ABBR_CODE" `
	Description string ` json:"description" `
	DateAdded   string ` json:"date_added" `
	AddedBy     string ` json:"added_by" `
	Status      string ` json:"status" `
}

//
func LoadAbbreviationsFromFile() (error) {

	bites, err := ioutil.ReadFile(abbrsFile)
	if err != nil {
		return err
	}
	var r abbrevReader
	err = json.Unmarshal(bites, &r)
	if err != nil {
		return err
	}

	a := new(Abbr)
	a.HeadCode = r.Info.Group
	//a.Class = r.Info.Class
	a.Description = r.Info.Heading

	a.Items = make([]AbbrevItem, len(r.Items))
	for i, h := range r.Items {
		aa := AbbrevItem{Item: h.Item, Description: h.Description,
			DateAdded: h.DateAdded, AddedBy: h.AddedBy, Status: h.Status}
		a.Items[i] = aa
	}
	//g.Index =
	return nil
}
