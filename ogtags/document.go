

package ogtags


import (
	//"encoding/csv"
	//"fmt"
	"strings"
	"encoding/csv"
	"io/ioutil"
	"path"
)

// The `Documen`t contains the data structure
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
	d.Init()
	return d
}

func (doc *Document) Init(){
	doc.Lines = make([]*Line, 0, 0)
	doc.GroupsIndex = make([]string, 0, 0)
	doc.Groups = make([]*GroupData, 0, 0)
}

func NewDocumentFromFile(file_path string) (*Document, error) {

	doc := NewDocument()
	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return doc, err
	}
	doc.FileName = path.Base(file_path)
	doc.Source = string(bites)
	err = doc.Parse()
	return doc, err
}

type Line struct {
	No 		int ` json:"no"  `
	Raw 	string  ` json:"raw"  `
	Columns []string  ` json:"columns"  `
	Errors 		[]string 	` json:"errors"  `
	Warnings 	[]string 	` json:"warnings"  `
}


func (doc *Document) Parse() error {

	doc.Init()
	// cleanup source, and split  into lines (unix style)
	raw_lines := strings.Split( strings.Replace(doc.Source, "\r", "", -1), "\n")

	// parse each csv_line into a Line object
	for idx, raw_line := range raw_lines {

		line := new(Line)
		line.No = idx + 1
		line.Raw = strings.TrimSpace(raw_line)

		// create records/columns for each line
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
		doc.Lines = append(doc.Lines, line)
	}


	// Map by "GROUP_CODE" to the data
	groups_map := make(map[string]*GroupData)

	// Pointer to current group
	var grp *GroupData

	// Current active group_code
	curr_group_code := ""


	// Walk though all the lines and extract group blocks
	for _, line := range doc.Lines {

		if line.Columns == nil {
			// ignore a blank line
			continue
		}

		col_count := len(line.Columns)

		// first column in the row type,
		switch line.Columns[0] {

		// The "GROUP","FOUR" = four character group name
		case GROUP:
			// A new groups, so set current object
			curr_group_code = line.Columns[1]
			// check group exists in map already...
			// this should always be ok,
			// TODO: possible errors =  double serialising groups
			_, found := groups_map[curr_group_code]
			if !found {
				// were now in this group
				grp = NewGroupData(curr_group_code)
				groups_map[curr_group_code] = grp
				doc.GroupsIndex = append(doc.GroupsIndex, curr_group_code)
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
				grp.Headings[c - 1].DType = line.Columns[c]
			}

		case UNIT:
			for c := 1; c < col_count; c++ {
				grp.Headings[c - 1].Unit = line.Columns[c]
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
	for _, g := range doc.GroupsIndex {
		doc.Groups = append(doc.Groups, groups_map[g])
	}


	return nil
}




