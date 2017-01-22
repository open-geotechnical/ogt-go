

package ags4


import (
	//"encoding/csv"
	//"fmt"
	"strings"
	"encoding/csv"

)

// The `Documen`t represents the data structure
// for an ags file
//
type Document struct {
	FileName string  	 ` json:"file_name"  `
	Source string   	 ` json:"source"  `
	Lines []*Line 		 ` json:"lines"  `
	GroupsIndex []string ` json:"groups_index" `
	Groups []*GroupData  ` json:"groups"  `
}

func NewDocument() *Document {
	d := new(Document)
	d.Lines = make([]*Line, 0, 0)
	return d
}

type Line struct {
	No 		int ` json:"no"  `
	Raw 	string  ` json:"raw"  `
	Columns []string  ` json:"columns"  `
	Errors 		[]string 	` json:"errors"  `
	Warnings 	[]string 	` json:"warnings"  `
}


func (this *Document) Parse() error {

	// cleanup source, and split  into lines (unix style)
	raw_lines := strings.Split( strings.Replace(this.Source, "\r", "", -1), "\n")

	// parse each csv_line into a Line object
	for idx, raw_line := range raw_lines {

		line := new(Line)
		line.No = idx + 1
		line.Raw = strings.TrimSpace(raw_line)

		// create records
		r := csv.NewReader(strings.NewReader(line.Raw))
		records, err := r.Read()
		if err != nil {
			//fmt.Println("err=", idx, err)
			// TODO, error EOF is fot blank lines, to need to ignore
			//
			line.Errors = append(line.Errors, err.Error())
		} else {
			line.Columns = records
		}
		this.Lines = append(this.Lines, line)
	}

	// Keep track of the groups "order" for serialisation
	groups_index_list := make([]string, 0, 0)

	// Map by "GROUP_CODE" to the data
	groups_map := make(map[string]*GroupData)

	// Pointer to current group
	var grp *GroupData

	// Current active group code
	curr_group_code := ""


	// Walk though all the lines
	for _, line := range this.Lines {

		if line.Columns == nil {
			// ignore a blank line
			continue
		}

		col_count := len(line.Columns)

		// Record[0] in first column in the row type,
		switch line.Columns[0] {

		// The "GROUP","FOUR" = four character group name
		case GROUP:

			curr_group_code = line.Columns[1]
			// check group exists in map already...
			// this should always be ok,
			// TODO: possible errors =  double serialising groups
			_, found := groups_map[curr_group_code]
			if !found {
				// were now in this group
				grp = NewGroupData(curr_group_code)
				groups_map[curr_group_code] = grp
				groups_index_list = append(groups_index_list, curr_group_code)
			} else {
				// OOPS, same groupname already exists wtf ??
				// recover with ??
			}

		// The "HEADING" is expected immedeately after the GROUP
		case HEADING:
			for c := 1; c < col_count; c++ {
				h := NewDataHeading(line.Columns[c])
				grp.Headings = append(grp.Headings, h)
			}

		case TYPE:
			for c := 1; c < col_count; c++ {
				grp.Headings[c - 1].SuggestedType = line.Columns[c]
			}

		case UNIT:
			for c := 1; c < col_count; c++ {
				grp.Headings[c - 1].SuggestedUnit = line.Columns[c]
			}

		case DATA:
			row := make(map[string]DataCell)
			for c := 1; c < col_count; c++ {
				hc := grp.Headings[c - 1].HeadCode
				row[hc] = DataCell{Value: line.Columns[c], HeadCode: hc, LineNo: line.No, ColNo: c}
				// TODO validate type and accuracy eg 2dp vs 3dp

			}
			grp.Data = append(grp.Data, row)
		}
	}

	// Serialise out in correct order
	for _, g := range groups_index_list {
		this.Groups = append(this.Groups, groups_map[g])
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




