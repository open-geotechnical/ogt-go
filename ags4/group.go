package ags4

import (
	"encoding/json"
	"errors"
	"fmt"
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


// Memory cache for classes and a filter
var Classes []string


// Memory cache for groups is in a map
var GroupsMap map[string]*Group

func init() {
	Classes = make([]string, 0, 0)
	GroupsMap = make(map[string]*Group)

}

func GetGroups() ([]*Group, error) {

	var keys []string
	for k := range GroupsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	groups := make([]*Group, 0, 0)
	for _, k := range keys {
		//fmt.Println("Key:", k, "Value:", m[k])
		groups = append(groups, GroupsMap[k])
	}

	return groups, nil
}

// Returns the group if found,
func GetGroup(group_code string) (*Group, error) {

	group_code = strings.TrimSpace(group_code)
	if len(group_code) != 4 {
		return nil, errors.New("Need four character group code")
	}
	group_code = strings.ToUpper(group_code)
	grp, ok := GroupsMap[group_code]
	if ok {
		return grp, nil
	}
	return nil, errors.New("Group '" + group_code + "' not found")
}

// Loads the groups.json file and bangs it into memory..
func LoadGroupsFromFile(file_path string)  error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	// need mutex here
	err = json.Unmarshal(bites, &GroupsMap)
	if err != nil {
		fmt.Println("errorf from json",  file_path, err)
		return err
	}
	fmt.Println("groups read from file ",  file_path)


	return nil
}
