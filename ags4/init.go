
package ags4

import (
	"fmt"
)

var DataDir string = ""

func InitLoad(data_dir string){

	DataDir = data_dir


	err := LoadUnitsFromFile(DataDir + "/units.json")
	if err != nil {
		fmt.Print(err)
	}

}

