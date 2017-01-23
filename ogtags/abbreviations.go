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
	AbbrsDDMap = make(map[string]*AbbrDD)
}

// AbbrsMap is the abbreviations mem variable
// with `head_code` pointer to abbr and items
var AbbrsDDMap map[string]*AbbrDD

// Represents an item in the abbreviations picklist
type AbbrDDItem struct {
	Code        string ` json:"code" `
	Description string ` json:"description" `
	//DateAdded   string ` json:"date_added" `
	//AddedBy     string ` json:"added_by" `
	List      string ` json:"list" `
}

// Represents an abbreviation picklist for the headcode, type PA
type AbbrDD struct {
	HeadCode    string       ` json:"head_code" `
	Picklist       []AbbrDDItem   ` json:"picklist" `
}

// Returns a list sorted hy headcode
func GetAbbrs() ([]*AbbrDD, error) {

	var keys []string
	for k := range AbbrsDDMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	abbrs := make([]*AbbrDD, 0, 0)
	for _, k := range keys {
		abbrs = append(abbrs, AbbrsDDMap[k])
	}
	return abbrs, nil
}

// Returns data on an abbreviation if found
func GetAbbr(head_code string) (*AbbrDD, bool, error) {

	head_code = strings.ToUpper(strings.TrimSpace(head_code))

	// TODO check case maybe and validate more ?
	// contains _, and 4 characters before
	if len(head_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	parts := strings.Split(head_code, "_")
	fmt.Println(parts)

	ab, found := AbbrsDDMap[head_code]
	return ab, found, nil
}

// Loads AGS abbreviations.json
func LoadAbbrsFromFile(file_path string) (error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	abbrsddMap  := make(map[string]*AbbrDD)
	err = json.Unmarshal(bites, &abbrsddMap)
	if err != nil {
		return err
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	AbbrsDDMap = abbrsddMap
	return nil
}
