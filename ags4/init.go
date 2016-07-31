
package ags4

import (
	"fmt"
)

// The directory where `ags-def-json/ags/4` is
// This is the cache variable and set with ags4.InitLoad()
var DataDir string = ""


// Ide here is to read all the def files..
// This needs to be a spperate function and reloadable with mutex
func InitLoad(data_dir string){


	DataDir = data_dir
	var all_errors []error

	err := LoadUnitsFromFile(DataDir + "/units.json")
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	}

	err = LoadGroupsIndexFromDir(DataDir + "/group/")
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	}

	if len(all_errors) > 0 {
		fmt.Println("WTF?", all_errors)
		// Tantrum()
	}


}

