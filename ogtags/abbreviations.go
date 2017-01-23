package ogtags

import (
	"encoding/json"
	"fmt"
	"errors"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)

func init() {
	AbbrsMap = make(map[string]*Abbr)
}

// AbbrsMap is the abbreviations mem_store/reader
// with head_code pointer to abbr and items
var AbbrsMap map[string]*Abbr


// Represents an item in the abbreviations picklist
type AbbrItem struct {
	Code        string ` json:"code" `
	Description string ` json:"description" `
	//DateAdded   string ` json:"date_added" `
	//AddedBy     string ` json:"added_by" `
	List      string ` json:"list" `
}

// Represents an abbreviation picklist for the headcode, type PA
type Abbr struct {
	HeadCode    string       ` json:"head_code" `
	Items       []AbbrItem   ` json:"abbrs" `
}

// Returns a list sorted hy headcode
func GetAbbrs() ([]*Abbr, error) {

	var keys []string
	for k := range AbbrsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	abbrs := make([]*Abbr, 0, 0)
	for _, k := range keys {
		abbrs = append(abbrs, AbbrsMap[k])
	}
	return abbrs, nil
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
func LoadAbbrsFromFile(file_path string) (error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	abbrsMap  := make(map[string]*Abbr)
	err = json.Unmarshal(bites, &abbrsMap)
	if err != nil {
		return err
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	AbbrsMap = abbrsMap
	return nil
}
