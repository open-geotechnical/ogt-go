package ogtags

import (
//"encoding/json"
//"io/ioutil"
//"sort"
//"fmt"
)

type Heading struct {
	HeadCode    string       ` json:"head_code" `
	HeadDescription string   ` json:"head_description" `
	DataType    string       ` json:"data_type" `
	Unit   string            ` json:"unit" `
	Example     string       ` json:"example" `
	RevDate     string       ` json:"rev_date" `
	SortOrder   int          ` json:"sort_order" `
	Status      string       ` json:"head_status" `
	//Picklist    []AbbrevItem ` json:"picklist" `
}

func (head *Heading)PickList() (*Abbr, bool, error) {
	return GetAbbr(head.HeadCode)
}


