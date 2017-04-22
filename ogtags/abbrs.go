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

// The AbbrsDataDictMap is lookup of the abbreviations in memory,
// and a lookup variable
// - The AGS data type is `PA` ie pick abbreviation
// - the map's key is the HEADING code eg SAMP_TYPE
// - pointing to the abbreviiation definition
// - Please create your own validation schema at will
var AbbrsDataDictMap map[string]*AbbrItem

func init() {
	AbbrsDataDictMap = make(map[string]*AbbrItem)
}

// Represents an Abbreviation item from a PA = pick abbr picklist)
// eg SAMP_TYPE = sample type
//   B = BUlk SAMPLE
//   CONC = Concrete Cube
//   W = Water
type AbbrItem struct {
	Code        string ` json:"abbr" `
	Description string ` json:"abbr_description" `

}

// Represents an abbreviations for the headcode, type PA
type AbbrDataDict struct {
	HeadCode    string       ` json:"head_code" `

	Abbrs       []AbbrItem   ` json:"abbrs" `
}

// Returns a list of abbreviations sorted by head_code
func AbbrsDataDictItems() ([]*AbbrDataDict, error) {

	var keys []string
	for k := range AbbrItem {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	abbrs := make([]*AbbrDataDict, 0, 0)
	for _, k := range keys {
		abbrs = append(abbrs, AbbrItem[k])
	}
	return abbrs, nil
}

// Returns data on an abbreviation if found
func AbbrDataDict(head_code string) (*AbbrDataDict, bool, error) {

	head_code = strings.ToUpper(strings.TrimSpace(head_code))

	// TODO check case maybe and validate more ?
	// contains _, and 4 characters before
	if len(head_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	parts := strings.Split(head_code, "_")
	fmt.Println(parts)

	ab, found := AbbrDataDictItem[head_code]
	return ab, found, nil
}

// Returns the picklicst for an abbrev
func AbbrDataDict(head_code string) (*AbbrDataDict, bool, error) {

	head_code = strings.ToUpper(strings.TrimSpace(head_code))

	// TODO check case maybe and validate more ?
	// contains _, and 4 characters before
	if len(head_code) < 6 {
		return nil, false, errors.New("Heading code too short")
	}
	parts := strings.Split(head_code, "_")
	fmt.Println(parts)

	ab, found := AbbrDataDictItem[head_code]
	return ab, found, nil
}

// Loads AGS abbreviations.json
func LoadAbbrsDDFromFile(file_path string) (error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	abbrsddMap  := make(map[string]*AbbrItem)
	err = json.Unmarshal(bites, &abbrsddMap)
	if err != nil {
		return err
	}
	var mutex = &sync.Mutex{}
	mutex.Lock()
	AbbrDataDictItem = abbrsddMap
	return nil
}
