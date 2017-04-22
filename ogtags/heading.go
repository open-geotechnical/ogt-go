package ogtags

import (
	//"encoding/json"
	//"io/ioutil"
	"errors"
	//"sort"
	//"fmt"
)


const (
	HEADING   = "HEADING" // the ags descriptor for a heading
)

var ErrHeadingInputInvalid = errors.New("Invalid heading code needs to be ABCD_XXX")
var ErrHeadingInvalidGroupCode = errors.New("Invalid heading code needs to be ABCD_XXX")
var ErrHeadingCodeNotInDataDict = errors.New("Heading code not in Data Dict")

// In ags the `Header` (think column) contains stuff about that column and its data
// The GROUP_HEADER in ags comprised the heading, type and unit..
//
type HeadingDataDict struct {
	// The
	HeadCode    string       ` json:"head_code" `
	HeadDescription string   ` json:"head_description" `
	DataType    string       ` json:"data_type" `
	Unit   string            ` json:"unit" `
	Example     string       ` json:"example" `
	RevDate     string       ` json:"rev_date" `
	SortOrder   int          ` json:"sort_order" `
	Status      string       ` json:"head_status" `
}


// Returns the abbreviation picklist for this headings.hHeadCode if found
func (head *HeadingDataDict)PickList() (*AbbrDD, bool, error) {
	return GetAbbrDD(head.HeadCode)
}


