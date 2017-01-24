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

// GroupDD is the data_dict of an agsgroup
type GroupDD struct {
	GroupCode   string    		`json:"group_code"`
	Class       string   		 `json:"class"`
	GroupDescription string    	`json:"group_description"`

	Parent string       `json:"parent"`
	Child string  `json:"child"`

	Headings    []HeadingDD		 `json:"headings"`
	Notes    []string		 `json:"notes"`
}

func init() {
	Classes = make([]string, 0, 0)
	GroupsDDMap = make(map[string]*GroupDD)
}

// Memory cache for classes and a filter
var Classes []string

// Memory cache for the Groups data dict is a map of four char group_code
var GroupsDDMap map[string]*GroupDD



// Returns the Groups data dict as a simple list/rows
func GetGroupsDD() ([]*GroupDD, error) {
	// first get key from map and sort
	var keys []string
	for k := range GroupsDDMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// construct list to return
	groups := make([]*GroupDD, 0, 0)
	for _, k := range keys {
		groups = append(groups, GroupsDDMap[k])
	}
	return groups, nil
}

// Returns the ags4 Group definition if found,
func GetGroupDD(group_code string) (*GroupDD, error) {

	// validate group_code
	group_code = strings.ToUpper(strings.TrimSpace(group_code))
	// TODO maybe a regex to ensure its 4 chars and upper case...maybe
	if len(group_code) != 4 {
		return nil, errors.New("Need four character  `group_code` ")
	}

	grp, found := GroupsDDMap[group_code]
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
	groupsMap := make(map[string]*GroupDD)
	err = json.Unmarshal(bites, &groupsMap)
	if err != nil {
		return err
	}

	// Success so update memory cache
	var mutex = &sync.Mutex{}
	mutex.Lock()
	GroupsDDMap = groupsMap
	mutex.Unlock()
	// TODO classes
	//for code, grp := range GroupsMap {

	//}
	return nil
}
