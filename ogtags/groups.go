package ogtags

import (
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"sort"
	"strings"
	"sync"
)
const (
	GROUP   = "GROUP" // the ags descriptor for a group
)

var ErrInvalidGroupCodeInput = errors.New("Invalid group code arg, must be four characters")
var ErrGroupCodeNotFound = errors.New("Group code not found in data dict")


// GroupDataDict is the data_dict of an ags4group
// from the data dictonary
type GroupDataDict struct {

	Class       string   		 `json:"class"`

	// The group CODE
	GroupCode   string    		`json:"group_code"`

	// The groups simple description
	GroupDescription string    	`json:"group_description"`

	// The parent group code
	Parent string       		`json:"parent"`
	// The child groups code
	Child string  				`json:"child"`

	Headings    []HeadingDataDict		 `json:"headings"`
	Notes    []string		 `json:"notes"`
>>>>>>> rename
}

func init() {
	Classes = make([]string, 0, 0)
<<<<<<< HEAD
	GroupsDDMap = make(map[string]*GroupDataDict)
=======
	GroupsDataDictMap = make(map[string]*GroupDataDict)
>>>>>>> rename
}

// Memory cache for the classes, uniquie and used for filter
var Classes []string

<<<<<<< HEAD
// Memory cache for the Groups data dict is a map of four char group_code
var GroupsDDMap map[string]*GroupDataDict



// Returns the Groups data dict as a simple list/rows
func GetGroupsDD() ([]*GroupDataDict, error) {
=======
// Memory cache of the four digit group code to a GroupDataDict
var GroupsDataDictMap map[string]*GroupDataDict



// Returns the Groups data dict as a list, sorted by group code
func GroupsDataDict() ([]*GroupDataDict, error) {
>>>>>>> rename
	// first get key from map and sort
	var keys []string
	for k := range GroupsDataDictMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// construct list to return
	groups := make([]*GroupDataDict, 0, 0)
	for _, k := range keys {
		groups = append(groups, GroupsDataDictMap[k])
	}
	return groups, nil
}

// Returns the Group Data Dict definition if found, else error
func GroupDataDict(group_code string) (*GroupDataDict, error) {

	// validate group_code
	group_code = strings.ToUpper(strings.TrimSpace(group_code))
	// TODO maybe a regex to ensure its 4 chars and upper case...maybe
	if len(group_code) != 4 {
		return nil, errors.New("Need four character  `group_code` ")
	}

	grp, found := GroupsDataDictMap[group_code]
	if found {
		return grp, nil
	}
	return nil, errors.New("Group code '" + group_code + "' not found")
}

// Loads the ags4 groups.json file from json into memory
func LoadGroupsDDFromFile(file_path string)  error {

	// read file
	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}

	// Read into memory
	groupsMap := make(map[string]*GroupDataDict)
	err = json.Unmarshal(bites, &groupsMap)
	if err != nil {
		return err
	}

	// Success so update memory cache
	var mutex = &sync.Mutex{}
	mutex.Lock()
	GroupsDataDictMap = groupsMap
	mutex.Unlock()
	// TODO classes
	//for code, grp := range GroupsMap {

	//}
	return nil
}
