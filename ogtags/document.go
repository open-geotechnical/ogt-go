
package ogtags

import (
	"strings"
	"encoding/csv"
	"io/ioutil"
	"path"
)

// The `Document` represents the data structure of geotechincal stuff
//
// It contains some meta data about the ags document,
// and parses the file into a mempry object for manipulation
// and output into various other formats, eg json, yaml, odf, email
//
// IMPORTANT:
// The ags file only contains the "data", other stuff such as images are
// contained within the Project
//
// ** Developer Goals **
//
// To make manipulation and validation of ags documents online
// and into workflow of
// - site pre research
// - and adding to stuff as we go along
//
// WIP:
// of an ags file, including the Raw source, lines and Groups of data
//
type Document struct {
	FileName string  	 ` json:"file_name"  `
	Source string   	 ` json:"source"  `
	Hash string          ` json:"hash"  `
	Lines []*Line 		 ` json:"lines"  `
	GroupsIndex []string ` json:"groups_index" `
	GroupsDataMap map[string]*GroupData  ` json:"groups"  `
}

// Create and initializes a Document pointer in memory
func NewDocument() *Document {
	doc := new(Document)
	doc.Lines = make([]*Line, 0, 0)
	doc.GroupsIndex = make([]string, 0, 0)
	doc.GroupsDataMap = make(map[string]*GroupData)
	return doc
}

// Returns a loaded Document object from a text ags file path, otherwise the err
// the loaded file could be an url or a file path,
// whicheved way doc will be loaded...
func LoadDocumentFromFile(url_path string) (*Document, error) {

	doc := NewDocument()
	bites, err := ioutil.ReadFile(url_path)
	if err != nil {
		return doc, err
	}
	doc.FileName = path.Base(url_path)
	doc.Source = string(bites)
	err = doc.Parse()
	return doc, err
}

// The Line object represents a line single from and ags file
// and is embedded within the Document
type Line struct {

	// The line no eg 1 for first (not zero index based)
	No 		int ` json:"no"  `

	// The raw string line as read, eg "GROUP","PROJ" stripped of white space ie no dangling \r or tabs
	Raw 	string  ` json:"raw"  `

	// The Columns is the result of golangs cvs parser PER line (cvs cant parse document as unequal column, etc)
	Columns []string  ` json:"columns"  `

	// Errors and warnings are WIP for indicating errors in file structure and data..
	Errors 		[]string 	` json:"errors"  `
	Warnings 	[]string 	` json:"warnings"  `
}


func (doc *Document) Parse() error {

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

	// Now we walk the lines, and extract groups and data

	// Pointer to current group
	var cgrp *GroupData

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
			_, found := doc.GroupsDataMap[curr_group_code]
			if !found {
				// were now in this group
				cgrp = NewGroupData(curr_group_code)
				doc.GroupsDataMap[curr_group_code] = cgrp
				doc.GroupsIndex = append(doc.GroupsIndex, curr_group_code)
			} else {
				// TODO, same groupname already exists wtf ??
				// recover with ??
			}

		// The "HEADING" is expected immedeately after the GROUP
		case HEADING:
			for c := 1; c < col_count; c++ {
				// TODO validate
				h := NewDataHeading(line.Columns[c])
				cgrp.Headings = append(cgrp.Headings, h)
			}

		case TYPE:
			for c := 1; c < col_count; c++ {
				// TODO validate
				cgrp.Headings[c - 1].DataType = line.Columns[c]
			}

		case UNIT:
			for c := 1; c < col_count; c++ {
				// TODO validate
				cgrp.Headings[c - 1].Unit = line.Columns[c]
			}

		case DATA:
			row := make(map[string]DataCell)
			for c := 1; c < col_count; c++ {
				hc := cgrp.Headings[c - 1].HeadCode
				row[hc] = DataCell{Value: line.Columns[c], HeadCode: hc, LineNo: line.No, ColIndex: c}
				// TODO validate type and accuracy eg 2dp vs 3dp

			}
			cgrp.Data = append(cgrp.Data, row)
		}
	}

	return nil
}




