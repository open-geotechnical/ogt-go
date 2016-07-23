
package ags4

import (
	"fmt"
)

var DataDir string = ""

func InitLoad(data_dir string){

	DataDir = data_dir
	var err error

	err = LoadUnitsFromFile(DataDir + "/units.json")
	if err != nil {
		fmt.Print(err)
	}

	err = LoadGroupsIndexFromFile(DataDir + "/groups.json")
	if err != nil {
		fmt.Print(err)
	}


}

