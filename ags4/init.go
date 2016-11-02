
package ags4

import (
	"fmt"
)

// The directory where `ags-def-json/ags/4` is
// This is the cache variable and set with ags4.InitLoad()
var DataDir string = ""
var groupsDir string = ""
var abbrevsDir string = ""
var examplesDir string = ""


// Ide here is to read all the def files..
// This needs to be a spperate function and reloadable with mutex
func InitLoad(data_dir string) {

	fmt.Println("Loading AGS")

	DataDir = data_dir
	abbrevsDir = DataDir + "/abbrev"
	groupsDir = DataDir + "/group"
	examplesDir = DataDir + "/examples"


	var all_errors []error

	err := LoadUnitsFromFile(DataDir + "/units.json")
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	}

	err = LoadAbbrevsFromDir()
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	}

	err = LoadGroupsFromDir(DataDir + "/group/")
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
		// Group errors go here
	}

	if len(all_errors) > 0 {
		fmt.Println("WTF?", all_errors)
		// Tantrum()
	}

	fmt.Println("Loading AGS DONE")


}

