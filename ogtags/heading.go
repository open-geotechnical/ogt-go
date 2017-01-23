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
	DType    string          ` json:"suggested_type" `
	Unit   string            ` json:"suggested_unit" `
	Example     string       ` json:"example" `
	RevDate     string       ` json:"rev_date" `
	SortOrder   int          ` json:"sort_order" `
	Status      string       ` json:"head_status" `
	//Picklist    []AbbrevItem ` json:"picklist" `
}


