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

// Groups data def contains and reads a Group
type GroupDD struct {
	GroupCode   string    		`json:"group_code"`
	Class       string   		 `json:"class"`
	GroupDescription string    	`json:"group_description"`

	Parent string       `json:"parent"`
	Child string  `json:"child"`

	Headings    []Heading		 `json:"headings"`
	Notes    []string		 `json:"notes"`
}

func init() {
	Classes = make([]string, 0, 0)
	GroupsDDMap = make(map[string]*GroupDD)
}

// Memory cache for classes and a filter
var Classes []string

// Memory cache for the groups is a map offour char  group_code
var GroupsDDMap map[string]*GroupDD



// Returns the ags4 groups sorted in a list/array
func GetGroups() ([]*GroupDD, error) {
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
func GetGroup(group_code string) (*GroupDD, error) {

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
func LoadGroupsFromFile(file_path string)  error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}

	// need mutex here
	groupsMap := make(map[string]*GroupDD)
	err = json.Unmarshal(bites, &groupsMap)
	if err != nil {
		return err
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	GroupsDDMap = groupsMap
	mutex.Unlock()
	// TODO classes
	//for code, grp := range GroupsMap {

	//}
	return nil
}
