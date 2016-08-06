package ags4

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

// AGS Group .. ta bill ..was Node
type Group struct {
	Code        string    `json:"code"`
	Class       string    `json:"class"`
	Description string    `json:"description"`
	Headings    []Heading `json:"headings"`
	//Units    	[]string   `json:"UNIT"`
	//Types    	[]string   `json:"TYPE"`
	//Data [][]string `json:"data"`
}


// Memory cache for groups
var Classes []string

// Memory cache for groups

var groupsMap map[string]*Group

func init() {
	groupsMap = make(map[string]*Group)
	Classes = make([]string, 0, 0)
}

func GetGroups() ([]*Group, error) {

	var keys []string
	for k := range groupsMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	groups := make([]*Group, 0, 0)
	for _, k := range keys {
		//fmt.Println("Key:", k, "Value:", m[k])
		groups = append(groups, groupsMap[k])
	}

	return groups, nil
}

func GetGroup(group_code string) (*Group, error) {

	group_code = strings.TrimSpace(group_code)
	if len(group_code) != 4 {
		return nil, errors.New("Need four letter group code")
	}
	group_code = strings.ToUpper(group_code)
	grp, ok := groupsMap[group_code]
	if ok {
		return grp, nil
	}
	return nil, errors.New("Group '" + group_code + "' not found")
}

// load groups.json file to mem
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
			groupsMap[grp.Code] = grp
		}
	}
	//fmt.Println(groupsMap)
	return nil
}

// Used to parse groups.json
type DEADgroupsReaderIndex struct {
	Meta []groupReaderInfo ` json:"meta" `
}

// Reader for group file (`groups/CODE.json`)
type groupFileReader struct {
	Info     groupReaderInfo      ` json:"info" `
	Notes    groupNotesReader     ` json:"notes" `
	Headings []groupHeadingReader ` json:"headings" `
}

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

func LoadGroupFromFile(file_path string) (*Group, error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return nil, err
	}
	var gr groupFileReader
	err = json.Unmarshal(bites, &gr)
	if err != nil {
		return nil, err
	}

	g := new(Group)
	g.Code = gr.Info.GroupCode
	g.Class = gr.Info.Class
	g.Description = gr.Info.Description

	g.Headings = make([]Heading, len(gr.Headings))
	for i, h := range gr.Headings {
		hh := Heading{Code: h.HeadCode, Description: h.Description,
			DataType: h.DataType, Unit: h.Unit, Status: h.Status,
			RevDate: h.RevDate, Sort: h.Sort, Example: h.Example}

		abbrs, found, erra := GetAbbrev(hh.Code)
		if erra != nil {
			fmt.Println("abbr not founc", hh.Code, erra)
		} else if found == true  {
			hh.Picklist = abbrs.Items
		}

		g.Headings[i] = hh
	}
	//g.Index =
	return g, nil
}
