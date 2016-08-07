package ags4

import (
	"encoding/json"
	//"fmt"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)

var abbrevsMap map[string]*Abbrev

func init() {
	abbrevsMap = make(map[string]*Abbrev)
	//Classes = make([]string, 0, 0)
}

type AbbrevItem struct {
	Item        string ` json:"item" `
	Description string ` json:"description" `
	DateAdded   string ` json:"date_added" `
	AddedBy     string ` json:"added_by" `
	Status      string ` json:"status" `
}

//
type Abbrev struct {
	HeadCode    string       ` json:"head_code" `
	Description string       ` json:"description" `
	Items       []AbbrevItem ` json:"items" `
}

func GetAbbrevs() ([]*Abbrev, error) {

	var keys []string
	for k := range abbrevsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	abbrevs := make([]*Abbrev, 0, 0)
	for _, k := range keys {
		//fmt.Println("Key:", k, "Value:", m[k])
		abbrevs = append(abbrevs, abbrevsMap[k])
	}
	return abbrevs, nil
}

func GetAbbrev(heading_code string) (*Abbrev, bool, error) {
	heading_code = strings.TrimSpace(heading_code)
	if len(heading_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	ab, found := abbrevsMap[heading_code]
	return ab, found, nil
}

/*
func DEADGetPicklist( heading_code string) ([]AbbrevItem, error) {
	heading_code = strings.TrimSpace(heading_code)
	if len(heading_code) < 6 {
		return nil, errors.New("Heading code too short")
	}
	ab, ok := abbrevsMap[heading_code]
	if !ok {
		return nil, errors.New("Heading code '" + heading_code + "' not found")
	}

	return ab.Items, nil
}
*/

// load groups.json file to mem
func LoadAbbrevsFromDir() error {

	//dir := DataDir + "/abbrev/"
	files, err := ioutil.ReadDir(abbrevsDir)
	if err != nil {
		return err
	}

	for _, f := range files {
		//fmt.Println("f=", f.Name())
		abr, errg := LoadAbbrevFromFile(f.Name())
		if errg != nil {
			//fmt.Println("err=", abr, errg)
		} else {
			//fmt.Println("ok=", abr.Heading)
			abbrevsMap[abr.HeadCode] = abr
		}
	}
	//fmt.Println(abbrevsMap)
	return nil
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
func LoadAbbrevFromFile(file_name string) (*Abbrev, error) {

	file_path := abbrevsDir + "/" + file_name
	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	var r abbrevReader
	err = json.Unmarshal(bites, &r)
	if err != nil {
		return nil, err
	}

	a := new(Abbrev)
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
	return a, nil
}
