package ogtags

import (
//"encoding/json"
//"io/ioutil"
	"errors"
//"sort"
//"fmt"
)

var ErrHeadingInputInvalid = errors.New("Invalid heading code needs to be ABCD_XXX")
var ErrHeadingInvalidGroupCode = errors.New("Invalid heading code needs to be ABCD_XXX")
var ErrHeadingCodeNotInDataDict = errors.New("Heading code not in Data Dict")

// The Heading DataDict container
type HeadingDataDict struct {
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


