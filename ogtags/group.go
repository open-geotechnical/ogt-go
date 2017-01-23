package ogtags

import (
	"encoding/json"
	"errors"
	//"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// AGS Group .
type Group struct {
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
	GroupsMap = make(map[string]*Group)
}

// Memory cache for classes and a filter
var Classes []string

// Memory cache for the groups is a map offour char  group_code
var GroupsMap map[string]*Group



// Returns the ags4 groups sorted in a list/array
func GetGroups() ([]*Group, error) {
	// first get key from map and sort
	var keys []string
	for k := range GroupsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	// construct list to return
	groups := make([]*Group, 0, 0)
	for _, k := range keys {
		groups = append(groups, GroupsMap[k])
	}

	return groups, nil
}

// Returns the ags4 Group definition if found,
func GetGroup(group_code string) (*Group, error) {

	// validate group_code
	group_code = strings.ToUpper(strings.TrimSpace(group_code))
	// TODO maybe a regex to ensure its 4 chars and upper case...maybe
	if len(group_code) != 4 {
		return nil, errors.New("Need four character  `group_code` ")
	}

	grp, found := GroupsMap[group_code]
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
	err = json.Unmarshal(bites, &GroupsMap)
	if err != nil {
		return err
	}
	return nil
}
