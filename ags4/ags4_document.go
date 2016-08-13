

package ags4


import (
	//"encoding/csv"
	"fmt"
	"strings"
	"encoding/csv"

)

// The Ags4Doc represents the data structure
// for an ags file
// The "struct should be
// cool and funky and serialiabled into json, yaml and AGS
// <NO xml.hell please ie no arrays said bill>
//

type Document struct {
	FileName string  	` json:"file_name"  `
	Source string   	` json:"source"  `
	Lines []*Line 		 ` json:"lines"  `
	//GroupsIndex map[string]*GroupData ` groups:"-"  `
	Groups []*GroupData ` json:"groups"  `
}

func NewDocument() *Document {
	d := new(Document)
	d.Lines = make([]*Line, 0, 0)
	return d
}

type Line struct {
	No 		int ` json:"no"  `
	Raw 	string  ` json:"raw"  `
	Records []string  ` json:"records"  `
	Errors 		[]string 	` json:"errors"  `
	Warnings 	[]string 	` json:"warnings"  `
}



func (this *Document) Parse() error {

	// Split raw string into Line objects
	lines := strings.Split(this.Source, "\r")

	for idx, raw_line := range lines {

		line := new(Line)
		line.No = idx + 1
		line.Raw = strings.TrimSpace(raw_line)
		r := csv.NewReader(strings.NewReader(line.Raw))
		record, err := r.Read()
		if err != nil {
			fmt.Println("err=", idx, err)
			line.Errors = append(line.Errors, err.Error())
		} else {
			line.Records = record
		}
		this.Lines = append(this.Lines, line)
	}

	gindex := make([]string, 0, 0)
	var grp *GroupData
	gmap := make(map[string]*GroupData)
	curr_group := ""
	//data_rows := make(map[int]map[string]DataCell)

	// Walk though all the lines
	for _, line := range this.Lines {
		if line.Records == nil {
			continue
		}

		// determine columns we need to loops
		col_count := len(line.Records)

		// Record[0] is first column in the row type,
		switch line.Records[0] {

		// The "GROUP","FOUR" = four character group name
		case GROUP:

			curr_group = line.Records[1]
			// check group exists in map already...
			// this should always be ok,
			// TODO: possible errors =  double serialising groups
			_, ok := gmap[curr_group]
			if !ok {
				// were now in this group
				grp = NewGroupData(curr_group)
				gmap[curr_group] = grp
				gindex = append(gindex, curr_group)
			} else {
				// OOPS, same groupname already exists wtf ??
				// recover with ??
			}

		// The "HEADING" is expetced after the GROUP
		case HEADING:
			for c := 1; c < col_count; c++ {
				h := NewDataHeading(line.Records[c])
				grp.Headings = append(grp.Headings, h)
			}

		case TYPE:
			for c := 1; c < col_count; c++ {
				grp.Headings[c - 1].DataType = line.Records[c]
			}

		case UNIT:
			for c := 1; c < col_count; c++ {
				grp.Headings[c - 1].Unit = line.Records[c]
			}

		case DATA:
			row := make(map[string]DataCell)
			for c := 1; c < col_count; c++ {
				hc := grp.Headings[c - 1].HeadCode
				row[hc] = DataCell{Value: line.Records[c], HeadCode: hc, LineNo: line.No, ColNo: c}
				// TODO validate type and accuracy eg 2dp vs 3dp

			}
			grp.Data = append(grp.Data, row)
		}
	}

	for _, g := range gindex {
		this.Groups = append(this.Groups, gmap[g])
	}

	/*
		//fmt.Println("l=", idx, line)
		r := csv.NewReader(strings.NewReader(line))
		record, err := r.Read()
		if err != nil {
			fmt.Println("err=", idx, err)
		} else {


			switch record[0] {
			case GROUP:
				curr_group := record[0]

				grp := NewGroupData(curr_group)
				//grpd.GroupCode = curr_group



				//group_codes = append(group_codes, curr_group)
				//group_lines[curr_group] = make([]string, 0, 0)

				//grp, ok := groupsMap[grp_code]
				//if !ok {
					// complain group not found
					//fmt.Println("err=", idx, "group not found")
				//} else {
					//grpd := new(GroupData)
					//grpd.GroupCode = grp.GroupCode
					//grpd.Description = grp.Description
					//grpd.Description = grp.Class
				//}



			//case DATA:
				//data = append(data, token)



			//case HEADING:
				//node.Heading = append(node.Heading, token)

			//case TYPE:
				//node.Type = append(node.Type, token)

			//case UNIT:
				//node.Unit = append(node.Unit, token)
			//default:
				//group_lines[curr_group] = append
			}

			//fmt.Println("ok=", idx, record[0], record[1])

		}
		*/
		/*
		for {
			record, err := r.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				fmt.Println("fatal", err)
				return err
			}

			fmt.Println("record=", len(record), record)
		}
		*/
	//}

	return nil
}




