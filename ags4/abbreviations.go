package ags4

import (
	"encoding/json"
	"fmt"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
)

func init() {
	AbbrsMap = make(map[string]*Abbr)
}

// AbbrsMap is the abbreviations mem_store/reader
// with head_code pointer to abbreviation and items
var AbbrsMap map[string]*Abbr


// Represents an item in the abbreviations picklist
type AbbrItem struct {
	Code        string ` json:"abbr_code" `
	Description string ` json:"abbr_desc" `
	//DateAdded   string ` json:"date_added" `
	//AddedBy     string ` json:"added_by" `
	List      string ` json:"abbr_list" `
}

// Represents an abbreviation picklist for the headcode, type PA
type Abbr struct {
	HeadCode    string       ` json:"head_code" `
	Items       []AbbrItem   ` json:"abbreviations" `
}

// Returns a list sorted
func GetAbbrs() ([]*Abbr, error) {

	var keys []string
	for k := range AbbrsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	abbrevs := make([]*Abbr, 0, 0)
	for _, k := range keys {
		//fmt.Println("Key:", k, "Value:", m[k])
		abbrevs = append(abbrevs, AbbrsMap[k])
	}
	return abbrevs, nil
}

// Returns data on an abbreviation if found
func GetAbbr(head_code string) (*Abbr, bool, error) {

	head_code = strings.ToUpper(strings.TrimSpace(head_code))

	// TODO check case maybe and validate more ?
	// contains _, and 4 characters before
	if len(head_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	parts := strings.Split(head_code, "_")
	fmt.Println(parts)
	ab, found := AbbrsMap[head_code]
	return ab, found, nil
}

// Loads AGS abbreviations.json
func LoadAbbreviationsFromFile(file_path string) (error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
    // TODO mutex
	err = json.Unmarshal(bites, &AbbrsMap)
	if err != nil {
		return err
	}
	return nil
}
