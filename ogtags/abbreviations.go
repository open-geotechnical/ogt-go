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

// AbbrsDDMap is the abbreviations data dict mem variable
// with `head_code` as key to abbr and picklist
var AbbrsDataDictMap map[string]*AbbrItem

func init() {
	AbbrsDataDictMap = make(map[string]*AbbrItem)
}

// Represents an Abbreviation item (a PA = pick abbr picklist)
// eg SAMP_TYPE = sample type
//   B =
//   CONC
//   W
type AbbrItem struct {
	Code        string ` json:"head_code" `
	Description string ` json:"description" `
	//DateAdded   string ` json:"date_added" `
	//AddedBy     string ` json:"added_by" `
	//List      string ` json:"list" `
}

// Represents an abbreviations for the headcode, type PA
type AbbrDataDict struct {
	HeadCode    string       ` json:"head_code" `
	Abbrs       []AbbrItem   ` json:"abbrs" `
}

// Returns a list of abbreviations sorted hy head_code
func AbbrsDataDict() ([]*AbbrDataDict, error) {

	var keys []string
	for k := range AbbrsDDMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	abbrs := make([]*AbbrDataDict, 0, 0)
	for _, k := range keys {
		abbrs = append(abbrs, AbbrsDDMap[k])
	}
	return abbrs, nil
}

// Returns data on an abbreviation if found
func GetAbbrDD(head_code string) (*AbbrDataDict, bool, error) {

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
func LoadAbbrsDDFromFile(file_path string) (error) {

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
