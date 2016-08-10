
package ags4

import (
	"strings"
)

type GroupData struct {
	GroupCode   string    `json:"group_code"`
	Class       string    `json:"class"`
	Description string    `json:"description"`

	// True if this in the Official Data dictionary
	AGSValid bool `json:"ags_valid"`

	Headings    []DataHeading ` json:"headings" `

	Data []map[string]DataCell ` json:"data" `
}


type DataHeading struct {
	Heading
	AGSValid bool `json:"ags_valid"`
	//Data []DataCell ` json:"data" `
}

type DataCell struct {
	HeadCode string   ` json:"head_code" `
	Value   string    ` json:"value" `
	Error   error    ` json:"error" `
	LineNo int  	` json:"line_no" `
	ColNo int  	` json:"col_no" `
}

//type DataRow struct {
//	Row
//}

func NewGroupData(grp_code string) *GroupData {

	gdata := new(GroupData)
	gdata.GroupCode = grp_code
	gdata.Headings = make([]DataHeading, 0, 0)

	grp_def, ok := groupsMap[gdata.GroupCode]
	if ok {
		gdata.Description = grp_def.Description
		gdata.Class = grp_def.Class
		gdata.AGSValid = true
	}
	return gdata
}


func NewDataHeading(head_code string) DataHeading {
	h := DataHeading{}
	h.HeadCode = head_code
	//h.Data = make([]DataCell, 0)

	parts := strings.Split(head_code, "_")
	grp, gok := groupsMap[parts[0]]
	if gok {
		for _, hd := range grp.Headings {
			if hd.HeadCode == head_code {
				h.AGSValid = true
				h.Description = hd.Description
				h.Picklist = hd.Picklist
				return h
			}
		}
	}
	return h
}