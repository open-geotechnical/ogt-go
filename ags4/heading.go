package ags4

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
	Unit        string       ` json:"unit" `
	Example     string       ` json:"example" `
	RevDate     string       ` json:"rev_date" `
	Sort        int          ` json:"sort" `
	Status      string       ` json:"status" `
	//Picklist    []AbbrevItem ` json:"picklist" `
}


