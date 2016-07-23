

package ags4

import (
	"encoding/json"
	"io/ioutil"
	//"fmt"
	"fmt"
)

// AGS Group .. ta bill ..was Node
type Group struct {
	GroupCode   string     `json:"GROUP"`
	Headings 	[]string   `json:"HEADING"`
	Units    	[]string   `json:"UNIT"`
	Types    	[]string   `json:"TYPE"`
	Data    	[][]string `json:"DATA"`
}

// Memory cache for groups
var Classes []string

// Memory cache for groups
var Groups []Group

func init(){
	Groups = make([]Group, 0, 0)
	Classes = make([]string, 0, 0)
}


// Used to parse groups.json
type groupsReaderIndex struct {
	Meta []groupReaderInfo ` json:"meta" `
}
type groupReaderInfo struct {
	GroupCode string	`json:"GROUP"`
	Class string	`json:"GROUP"` // ignored as in group file
	Description string	`json:"GROUP"`  // ignored as in group file
}

// load groups.json file to mem
func LoadGroupsIndexFromFile(file_path string) error {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	var gr groupsReaderIndex
	err = json.Unmarshal(bites, &gr)
	if err != nil {
		return err
	}

	for _, g := range gr.Meta {
		fmt.Println(g)
		grp, e := LoadGroupFromFile(DataDir + "/groups/" + g.GroupCode + ".json")
		if e != nil {
			// TODO log
		} else{
			Groups = append(Groups, grp)
		}
	}


	return nil
}

type groupNotesReader string {
	List []string ` json:"list" `
}

type groupFileReader struct {
	Info groupReaderInfo 	` json:"info" `
	Notes groupNotesReader 	` json:"notes" `

}


func LoadGroupFromFile(file_path string) (Group, error) {

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}
	var gr groupsIndexReader
	err = json.Unmarshal(bites, &gr)
	if err != nil {
		return err
	}

}