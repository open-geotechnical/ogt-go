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
		return nil, errors.New("Need four letter group code")
	}
	group_code = strings.ToUpper(group_code)
	grp, ok := GroupsMap[group_code]
	if ok {
		return grp, nil
	}
	return nil, errors.New("Group '" + group_code + "' not found")
}

// load groups.json file to mem
/*
func LoadGroupsFromDir(groups_dir string) error {

	files, err := ioutil.ReadDir(groups_dir)
	if err != nil {
		return err
	}

	for _, f := range files {
		//fmt.Println("f=", f.Name())
		grp, errg := LoadGroupFromFile(groups_dir + "/" + f.Name())
		if errg != nil {
			//fmt.Println("err=", grp, errg)
		} else {
			//fmt.Println("ok=", grp.Code)
			groupsMap[grp.GroupCode] = grp
		}
	}
	//fmt.Println(groupsMap)
	return nil
}
*/



/*
type groupReaderInfo struct {
	GroupCode   string `json:"GROUP"`
	Class       string `json:"class"`       // ignored as in group file
	Description string `json:"description"` // ignored as in group file
}

type groupHeadingReader struct {
	HeadCode    string ` json:"HEADING" `
	Description string ` json:"description" `
	DataType    string ` json:"data_type" `
	Unit        string ` json:"unit" `
	Example     string ` json:"example" `
	RevDate     string ` json:"rev_date" `
	Sort        int    ` json:"sort" `
	Status      string ` json:"status" `
}
type groupNotesReader struct {
	List []string ` json:"list" `
}
*/
// Reader for group file (`groups/CODE.json`)
/*
type groupsFileReader struct {
	groupReaderInfo      ` json:"info" `
	Notes    groupNotesReader     ` json:"notes" `
	Headings []groupHeadingReader ` json:"headings" `
}
*/
func LoadGroupsFromFile(file_path string)  error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	var grps map[string]*Group
	err = json.Unmarshal(bites, &grps)
	if err != nil {
		fmt.Println("errorf from json",  file_path, err)
		return err
	}
	fmt.Println("groups read from json",  file_path)
	// So this is the dynamic update
	// need mutex here
	GroupsMap = grps

	return nil
}
