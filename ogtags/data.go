
package ogtags

import (
	"strings"
)


type GroupData struct {

	Class       string    	` json:"class" 		yaml:"class" `
	GroupCode   string    	` json:"group_code" yaml:"group_code" `
	GroupDescription string	` json:"group_description" yaml:"group_description" `

	Headings    []DataHeading ` json:"headings" yaml:"headings" `
	Data []map[string]DataCell ` json:"data" yaml:"data" `
}


type DataHeading struct {
	HeadingDD
	Valid bool `json:"valid"`
}

// The data cell represent an indivudual Cell, or row column of data,
// The DataCell has its code and atrributes
// The "FIELD" must also comform to validation
// This might mean achecking values against one anothere.
// or labeling samples, and inputs outta sequence..
// BUt ?? whats your input ??
//
type DataCell struct {
	HeadCode string   ` json:"head_code" `
	Value   string    ` json:"value" `
	Error   error    ` json:"error" `
	LineNo int  	` json:"line_no" `
	ColIndex int  	` json:"col_index" `
}

func NewGroupData(grp_code string) *GroupData {

	gdata := new(GroupData)
	gdata.GroupCode = grp_code
	gdata.Headings = make([]DataHeading, 0, 0)

	// Check the definition exists and use it
	grpdd, ok := GroupsDataDictMap[gdata.GroupCode]
	if ok {
		gdata.GroupDescription = grpdd.GroupDescription
		gdata.Class = grpdd.Class
		gdata.Valid = true
	}
	return gdata
}


func NewDataHeading(head_code string) DataHeading {
	h := DataHeading{}
	h.HeadCode = head_code
	//h.Data = make([]DataCell, 0)

	parts := strings.Split(head_code, "_")
	grp, gok := GroupsDataDictMap[parts[0]]
	if gok {
		for _, hd := range grp.Headings {
			if hd.HeadCode == head_code {
				h.Valid = true
				//h.Description = hd.Description
				//h.Picklist = hd.Picklist
				return h
			}
		}
	}
	return h
}