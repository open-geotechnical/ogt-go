package ogtags

import (
//"encoding/json"
//"io/ioutil"
//"sort"
//"fmt"
)

// The Heading DataDict container
type HeadDataDict struct {
	HeadCode    string       ` json:"head_code" `
	HeadDescription string   ` json:"head_description" `
	DataType    string       ` json:"data_type" `
	Unit   string            ` json:"unit" `
	Example     string       ` json:"example" `
	RevDate     string       ` json:"rev_date" `
	SortOrder   int          ` json:"sort_order" `
	Status      string       ` json:"head_status" `
}

// Returns a picklist for this headings.hHeadCode if found
func (head *HeadingDataDict)PickList() (*AbbrDD, bool, error) {
	return GetAbbrDD(head.HeadDataDict)
}


