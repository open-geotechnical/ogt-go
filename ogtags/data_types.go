package ogtags

import (
	"io/ioutil"
	"encoding/json"
	"sync"
)

// The DataType represents the Headings Type expected,
// and is defined in AGS see
type DataType struct {
	Code string
	Description string
}

var DataTypesDDMap  map[string]*DataTypeDD

func init(){
	DataTypesDDMap = make(map[string]*DataTypeDD)
}

func LoadDataTypesDDFromFile(file_path string) error{

	bites, err := ioutil.ReadFile(file_path)
	if err != nil {
		return err
	}

	dataTypesDDMap := make(map[string]*DataTypeDD)
	err = json.Unmarshal(bites, &dataTypesDDMap)
	if err != nil {
		return err
	}

	var mutex = &sync.Mutex{}
	mutex.Lock()
	DataTypesDDMap = dataTypesDDMap
	mutex.Unlock()

	return nil
}