
package ogtags

import (
	"fmt"
)

// The directory where `ags-data-dict/ags4` is
// This is the cache variable and set with ags4.InitLoad()
var DataDictPath string = ""
var groupsFile string = ""
var abbrsFile string = ""
var examplesDir string = ""


// Ide here is to read all the def files..
// This needs to be a spperate function and reloadable with mutex
func InitLoad(ags_data_dict_dir string) {

	fmt.Println("Loading AGS", ags_data_dict_dir)

	DataDictPath = ags_data_dict_dir
	groupsFile = DataDictPath + "/ags4/groups.json"
	abbrsFile = DataDictPath + "/ags4/abbreviations.json"

	examplesDir = DataDictPath + "/ags4_examples"


	var all_errors []error
	var err error

	// TODO said pedro
	/*
	err := LoadUnitsFromFile(DataDictPath + "/units.json")
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	} */

	// groups and their heading need to be loaded first
	err = LoadGroupsFromFile(groupsFile)
	if err != nil {
		fmt.Print(err)

		all_errors = append(all_errors, err)
		// Group errors go here
		return
	} else {
		fmt.Println("loaded groups")
	}

	err = LoadAbbrsFromFile(abbrsFile)
	if err != nil {
		fmt.Print(err)
		all_errors = append(all_errors, err)
	}



	if len(all_errors) > 0 {
		fmt.Println("WTF?", all_errors)
		// Tantrum()
	}

	fmt.Println("Loading AGS DONE")


}

